package moac.ipfs.modules.web.service.impl;


import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.baomidou.mybatisplus.mapper.EntityWrapper;
import com.baomidou.mybatisplus.service.impl.ServiceImpl;
import moac.ipfs.common.exception.RRException;
import moac.ipfs.common.httputils.OkHttpClients;
import moac.ipfs.common.httputils.OkHttpParam;
import moac.ipfs.common.httputils.OkhttpResult;
import moac.ipfs.common.utils.*;
import moac.ipfs.modules.back.subchain.entity.SubchainEntity;
import moac.ipfs.modules.back.subchain.service.SubchainService;
import moac.ipfs.modules.back.user.entity.FileLogEntity;
import moac.ipfs.modules.back.user.entity.UserFileEntity;
import moac.ipfs.modules.back.user.service.FileLogService;
import moac.ipfs.modules.back.user.service.UserFileService;
import moac.ipfs.modules.back.user.service.UserService;
import moac.ipfs.modules.web.form.ImportAddressForm;
import moac.ipfs.modules.web.form.CreateAddressForm;
import moac.ipfs.modules.back.user.dao.UserDao;
import moac.ipfs.modules.back.user.entity.UserEntity;
import moac.ipfs.modules.web.form.FileParamsForm;
import moac.ipfs.modules.web.form.QueryFileListForm;
import moac.ipfs.modules.web.service.BaseWebService;
import org.apache.commons.lang.StringUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import javax.servlet.http.HttpServletRequest;
import java.math.BigDecimal;
import java.util.*;


/**
 * @author GZC
 */
@Service("baseWebService")
public class BaseWebServiceImpl extends ServiceImpl<UserDao, UserEntity> implements BaseWebService {
    Logger logger = LoggerFactory.getLogger(BaseWebServiceImpl.class);

    @Autowired
    private UserService userService;
    @Autowired
    private UserFileService userFileService;
    @Autowired
    private SubchainService subchainService;
    @Autowired
    private RedisUtils redisUtils;
    @Autowired
    private FileLogService fileLogService;

    @Override
    @Transactional(rollbackFor = Exception.class)
    public UserEntity importAddressService(ImportAddressForm form, HttpServletRequest request) {
        OkHttpParam okHttpParam = new OkHttpParam();
        okHttpParam.setApiUrl(Constant.IMPORT_ADDRESS_URL);
        Map<String, String> map = new HashMap<>(16);
        map.put("importType", form.getImportType());
        map.put("keyStore", form.getKeyStore() == null ? "" : form.getKeyStore());
        map.put("password", form.getPassword());
        map.put("mnemonic", form.getMnemonic() == null ? "" : form.getMnemonic());
        map.put("encryption", form.getEncryption() == null ? "" : form.getEncryption());
        map.put("plaintextPrivateKey", form.getPlaintextPrivateKey() == null ? "" : form.getPlaintextPrivateKey());
        OkhttpResult result = null;
        JSONObject json = null;
        try {
            result = OkHttpClients.post(okHttpParam, map, String.class, OkHttpClients.SYNCHRONIZE);
            json = JSON.parseObject(result.getResult());
            if (!"success".equals(json.getString("message"))) {
                logger.error(json.toJSONString());
                throw new RRException("请求导入地址错误!");
            }
            JSONObject jsonObject = json.getJSONObject("resultData");
            UserEntity userEntity = userService.selectOne(new EntityWrapper<UserEntity>().eq("address", jsonObject.getString("address")));
            if (userEntity != null) {
                userEntity.setLastLoginTime(System.currentTimeMillis());
                userEntity.setLastLoginIp(IPUtils.getIpAddr(request));
                if (Constant.PLAINTEXTPRIVATEKEY_TYPE.equals(form.getImportType())) {
                    userEntity.setPassword(form.getPassword());
                }
                this.updateById(userEntity);
            } else {
                userEntity = new UserEntity();
                userEntity.setCreateTime(System.currentTimeMillis());
                userEntity.setLastLoginIp(IPUtils.getIpAddr(request));
                userEntity.setPassword(form.getPassword());
                if (StringUtils.isNotBlank(form.getPasswordHint())) {
                    userEntity.setPasswordHint(form.getPasswordHint());
                }
                userEntity.setUserAccount("IPFS" + Utils.getSerialNo());
                userEntity.setAddress(jsonObject.getString("address"));
                userEntity.setKeyStore(jsonObject.getString("keyStore"));
                this.insert(userEntity);
            }
            return userEntity;
        } catch (Exception e) {
            e.printStackTrace();
            throw new RRException("请求导入地址错误!");
        }
    }

    @Override
    @Transactional(rollbackFor = Exception.class)
    public UserEntity createAddressService(CreateAddressForm form, HttpServletRequest request) {
        OkHttpParam okHttpParam = new OkHttpParam();
        okHttpParam.setApiUrl(Constant.CREATE_ADDRESS_URL);
        Map<String, String> map = new HashMap<>(16);
        map.put("password", form.getPassword());
        OkhttpResult result = null;
        JSONObject json = null;
        try {
            result = OkHttpClients.post(okHttpParam, map, String.class, OkHttpClients.SYNCHRONIZE);
            json = JSON.parseObject(result.getResult());
            if (!"success".equals(json.getString("message"))) {
                logger.error(json.toJSONString());
                throw new RRException("请求创建地址错误!");
            }
        } catch (Exception e) {
            e.printStackTrace();
            throw new RRException("请求创建地址错误!");
        }
        JSONObject jsonObject = json.getJSONObject("resultData");
        UserEntity userEntity = new UserEntity();
        userEntity.setCreateTime(System.currentTimeMillis());
        userEntity.setLastLoginIp(IPUtils.getIpAddr(request));
        userEntity.setPassword(form.getPassword());
        userEntity.setLastLoginTime(System.currentTimeMillis());
        if (StringUtils.isNotBlank(form.getPasswordHint())) {
            userEntity.setPasswordHint(form.getPasswordHint());
        }
        userEntity.setUserAccount("IPFS" + Utils.getSerialNo());
        userEntity.setAddress(jsonObject.getString("address"));
        userEntity.setKeyStore(jsonObject.getString("keystore"));
        this.insert(userEntity);
        return userEntity;
    }

    @Override
    @Transactional(rollbackFor = Exception.class)
    public void addFileService(FileParamsForm form) {
        UserEntity userEntity = userService.selectOne(new EntityWrapper<UserEntity>().eq("address",form.getAddress()));
        UserFileEntity userFileEntity = userFileService.selectOne(new EntityWrapper<UserFileEntity>().eq("hash", form.getHash()).eq("user_id",userEntity.getUserId()));
        if (userFileEntity != null) {
            throw new RRException("文件已存在");
        }
        OkHttpParam okHttpParam = new OkHttpParam();
        okHttpParam.setApiUrl(Constant.SAVEFILE_URL);
        Map<String, String> map = new HashMap<>(16);
        map.put("address", form.getAddress());
        map.put("filePath", form.getFilePath());
        map.put("fileName", form.getFileName());
        map.put("fileRealSize", form.getFileRealSize());
        map.put("password",userEntity.getPassword());
        map.put("keyStore",userEntity.getKeyStore());
        /**
         * 新增文件表记录
         */
        userFileEntity = new UserFileEntity();
        userFileEntity.setCreateTime(System.currentTimeMillis());
        userFileEntity.setFileRealSize(new BigDecimal(form.getFileRealSize()));
        userFileEntity.setFileSize(Utils.getDataSize(Long.valueOf(form.getFileRealSize())));
        userFileEntity.setRemark(form.getRemark());
        userFileEntity.setUserId(userService.queryUserIdByAddress(form.getAddress()));
        userFileEntity.setFileName(form.getFileName());
        userFileEntity.setEncrypt(form.getEncrypt());
        userFileService.insert(userFileEntity);
        map.put("createTime",String.valueOf(userFileEntity.getCreateTime()));
        map.put("encrypt", String.valueOf(form.getEncrypt()));
        map.put("fileId", String.valueOf(userFileEntity.getFileId()));
        SubchainEntity subchainEntity = selectSubchainAddress(form.getFileRealSize());
        /**
         * 更新存储子链信息
         */
        subchainEntity.setRemainSize(subchainEntity.getRemainSize().subtract(new BigDecimal(form.getFileRealSize())));
        subchainEntity.setUseSize(subchainEntity.getUseSize().add(new BigDecimal(form.getFileRealSize())));
        subchainEntity.setPercentageUse(subchainEntity.getUseSize().divide(subchainEntity.getSubchainSize(), 2, BigDecimal.ROUND_HALF_UP));
        subchainService.updateById(subchainEntity);
        map.put("subchainAddress", subchainEntity.getSubchainAddress());
        map.put("subchainSize", subchainEntity.getSubchainSize().toPlainString());
        map.put("remainSize", subchainEntity.getRemainSize().toPlainString());
        BigDecimal percentageUse = subchainEntity.getPercentageUse().multiply(new BigDecimal("1000000"));
        map.put("percentageUse", percentageUse.stripTrailingZeros().toPlainString());
        OkhttpResult result = null;
        JSONObject json = null;
        try {
            result = OkHttpClients.post(okHttpParam, map, String.class, OkHttpClients.SYNCHRONIZE);
            json = JSON.parseObject(result.getResult());
            JSONObject jsonObject = json.getJSONObject("resultData");
            if (!"success".equals(json.getString("message"))) {
                logger.error(json.toJSONString());
                throw new RRException("新增文件错误!");
            }
            userFileEntity.setHash(jsonObject.getString("hash"));
            userFileEntity.setSubchainAddress(subchainEntity.getSubchainAddress());
            userFileService.updateById(userFileEntity);
            /**
             * 新增文件操作日志
             */
            doFileLog(userFileEntity,form.getAddress(),"上传文件");
        } catch (Exception e) {
            e.printStackTrace();
            throw new RRException("新增文件错误!");
        }
    }

    @Override
    public String queryFileByHashService(FileParamsForm form) {
        OkHttpParam okHttpParam = new OkHttpParam();
        okHttpParam.setApiUrl(Constant.READFILE_URL);
        Map<String, String> map = new HashMap<>(16);
        map.put("address", form.getAddress());
        map.put("hash", form.getHash());
        UserFileEntity userFileEntity = userFileService.selectOne(new EntityWrapper<UserFileEntity>().eq("hash", form.getHash()));
        map.put("subchainAddress", "0x5Fc8aE6BB3CfBb91d0617A58Bb115f991f245020");
        map.put("fileName",userFileEntity.getFileName());
        map.put("encrypt", String.valueOf(userFileEntity.getEncrypt()));
        UserEntity userEntity = userService.selectOne(new EntityWrapper<UserEntity>().eq("address",form.getAddress()));
        map.put("password",userEntity.getPassword());
        map.put("keyStore",userEntity.getKeyStore());
        OkhttpResult result = null;
        JSONObject json = null;
        try {
            result = OkHttpClients.post(okHttpParam, map, String.class, OkHttpClients.SYNCHRONIZE);
            json = JSON.parseObject(result.getResult());
            if (!"success".equals(json.getString("message"))) {
                logger.error(json.toJSONString());
                throw new RRException("查询文件错误!");
            }
        } catch (Exception e) {
            e.printStackTrace();
            throw new RRException("查询文件错误!");
        }
        JSONObject jsonObject = json.getJSONObject("resultData");
        return jsonObject.getString("fileUrl");
    }

    @Override
    @Transactional(rollbackFor = Exception.class)
    public void deleteFileService(FileParamsForm form) {
        UserFileEntity userFileEntity = userFileService.selectOne(new EntityWrapper<UserFileEntity>().eq("hash", form.getHash()));
        if (userFileEntity.getDeleteType() == 1){
            throw new RRException("文件已删除!");
        }
        OkHttpParam okHttpParam = new OkHttpParam();
        okHttpParam.setApiUrl(Constant.REMOVEFILE_URL);
        Map<String, String> map = new HashMap<>(16);
        map.put("address", form.getAddress());
        map.put("hash", form.getHash());
        map.put("subchainAddress", userFileEntity.getSubchainAddress());
        map.put("fileRealSize", userFileEntity.getFileRealSize().toPlainString());
        map.put("fileId",String.valueOf(userFileEntity.getFileId()));
        map.put("createTime",String.valueOf(userFileEntity.getCreateTime()));
        map.put("fileName",userFileEntity.getFileName());
        UserEntity userEntity = userService.selectOne(new EntityWrapper<UserEntity>().eq("address",form.getAddress()));
        map.put("password",userEntity.getPassword());
        map.put("keyStore",userEntity.getKeyStore());
        /**
         * 更新存储子链信息
         */
        SubchainEntity subchainEntity = subchainService.selectOne(new EntityWrapper<SubchainEntity>().eq("subchain_address", userFileEntity.getSubchainAddress()));
        subchainEntity.setRemainSize(subchainEntity.getRemainSize().add(userFileEntity.getFileRealSize()));
        subchainEntity.setUseSize(subchainEntity.getUseSize().subtract(userFileEntity.getFileRealSize()));
        subchainEntity.setPercentageUse(subchainEntity.getUseSize().divide(subchainEntity.getSubchainSize(), 2, BigDecimal.ROUND_HALF_UP));
        subchainService.updateById(subchainEntity);
        map.put("subchainAddress", subchainEntity.getSubchainAddress());
        map.put("subchainSize", subchainEntity.getSubchainSize().toPlainString());
        map.put("remainSize", subchainEntity.getRemainSize().toPlainString());
        BigDecimal percentageUse = subchainEntity.getPercentageUse().multiply(new BigDecimal("1000000"));
        map.put("percentageUse", percentageUse.stripTrailingZeros().toPlainString());
        OkhttpResult result = null;
        JSONObject json = null;
        try {
            result = OkHttpClients.post(okHttpParam, map, String.class, OkHttpClients.SYNCHRONIZE);
            json = JSON.parseObject(result.getResult());
            if ("success".equals(json.getString("message"))) {
                /**
                 * 更新文件逻辑删除
                 */
                userFileEntity.setDeleteType(1);
                userFileService.updateById(userFileEntity);
                /**
                 * 新增文件操作日志
                 */
                doFileLog(userFileEntity,form.getAddress(),"删除文件");
            } else {
                logger.error(json.toJSONString());
                throw new RRException("请求删除文件错误!");
            }
        } catch (Exception e) {
            e.printStackTrace();
            throw new RRException("请求删除文件错误!");
        }
    }

    //TODO 打包注意
    public SubchainEntity selectSubchainAddress(String fileSize) {
//        List<String> recentlyUsedList = JSON.parseArray(redisUtils.get("recentlyUsed"), String.class);
        Map<String, Object> params = new HashMap<>(16);
//        if (recentlyUsedList != null){
//            params.put("list", recentlyUsedList);
//        }
        params.put("fileSize", fileSize);
        SubchainEntity subchainAddress = subchainService.queryAddFileList(params);
//        recentlyUsedList.add(subchainAddress.getSubchainAddress());
//        if (recentlyUsedList.size() > 10){
//            recentlyUsedList.remove(0);
//        }
//        String str = JSON.toJSONString(recentlyUsedList);
//        redisUtils.set("recentlyUsed", str, -1);
        return subchainAddress;
    }

    /**
     * 新增文件操作日志
     * @param userFileEntity
     * @param address
     * @param operationType
     */
    private void doFileLog(UserFileEntity userFileEntity, String address, String operationType){
        FileLogEntity fileLogEntity = new FileLogEntity();
        fileLogEntity.setFileHash(userFileEntity.getHash());
        fileLogEntity.setFileRemark(userFileEntity.getRemark());
        fileLogEntity.setOperationTime(System.currentTimeMillis());
        fileLogEntity.setOperationType(operationType);
        fileLogEntity.setUserAddress(address);
        fileLogEntity.setUserId(userFileEntity.getUserId());
        fileLogService.insert(fileLogEntity);
    }
}