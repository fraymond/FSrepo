package moac.ipfs;


import com.alibaba.fastjson.JSON;
import com.baomidou.mybatisplus.mapper.EntityWrapper;
import moac.ipfs.common.utils.RedisUtils;
import moac.ipfs.common.utils.Utils;
import moac.ipfs.modules.back.currency.entity.CurrencyEntity;
import moac.ipfs.modules.back.currency.service.CurrencyService;
import moac.ipfs.modules.back.storagePackage.entity.StoragePackageEntity;
import moac.ipfs.modules.back.storagePackage.service.StoragePackageService;
import moac.ipfs.modules.back.subchain.dao.SubchainDao;
import moac.ipfs.modules.back.subchain.entity.SubchainEntity;
import moac.ipfs.modules.back.subchain.entity.SubchainIpEntity;
import moac.ipfs.modules.back.subchain.service.SubchainIpService;
import moac.ipfs.modules.back.subchain.service.SubchainService;
import moac.ipfs.modules.back.sys.entity.SysUserEntity;
import moac.ipfs.modules.back.user.entity.UserAssetsEntity;
import moac.ipfs.modules.back.user.entity.UserEntity;
import moac.ipfs.modules.back.user.service.UserAssetsService;
import moac.ipfs.modules.back.user.service.UserService;
import moac.ipfs.modules.web.service.impl.BaseWebServiceImpl;
import moac.ipfs.service.DataSourceTestService;
import org.apache.commons.lang.builder.ToStringBuilder;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import java.math.BigDecimal;
import java.util.*;


@RunWith(SpringRunner.class)
@SpringBootTest
public class DynamicDataSourceTest {
    @Autowired
    private DataSourceTestService dataSourceTestService;
    @Autowired
    private BaseWebServiceImpl baseWebService;
    @Autowired
    private RedisUtils redisUtils;
    @Autowired
    private SubchainService subchainService;
    @Autowired
    private SubchainIpService subchainIpService;
    @Autowired
    private SubchainDao subchainDao;
    @Autowired
    private UserService userService;
    @Autowired
    private UserAssetsService userAssetsService;
    @Autowired
    private CurrencyService currencyService;
    @Autowired
    private StoragePackageService storagePackageService;

    @Test
    public void test() {
        //数据源1
        SysUserEntity user1 = dataSourceTestService.queryUser(1L);
        System.out.println(ToStringBuilder.reflectionToString(user1));

        //数据源2
        SysUserEntity user2 = dataSourceTestService.queryUser2(1L);
        System.out.println(ToStringBuilder.reflectionToString(user2));

        //数据源1
        SysUserEntity user3 = dataSourceTestService.queryUser(1L);
        System.out.println(ToStringBuilder.reflectionToString(user3));
    }

    @Test
    public void test2() {
        Map<String, Object> params1 = new HashMap<>(16);
        params1.put("fileSize", 10000);
        SubchainEntity subchainEntity = subchainService.queryAddFileList(params1);
        List<String> recentlyUsedList = new ArrayList<>();
        recentlyUsedList.add(subchainEntity.getSubchainAddress());
        String str = JSON.toJSONString(recentlyUsedList);
        redisUtils.set("recentlyUsed", str, -1);
        List<String> recentlyUsed = JSON.parseArray(redisUtils.get("recentlyUsed"), String.class);
        Map<String, Object> params = new HashMap<>(16);
        if (recentlyUsed != null) {
            params.put("list", recentlyUsed);
        }
        params.put("fileSize", 10000);
        SubchainEntity subchainAddress = subchainService.queryAddFileList(params);
        recentlyUsed.remove(0);
        recentlyUsed.add(subchainAddress.getSubchainAddress());
        String strr = JSON.toJSONString(recentlyUsed);
        redisUtils.set("recentlyUsed", strr, -1);
        System.out.println(strr);
    }

    @Test
    public void test3() {
        Map<String, Object> params1 = new HashMap<>(16);
        params1.put("fileSize", 10000);
        SubchainEntity subchainEntity = subchainService.queryAddFileList(params1);
        BigDecimal percentageUse = subchainEntity.getPercentageUse().multiply(new BigDecimal("1000000"));
        String str = percentageUse.stripTrailingZeros().toPlainString();
        System.out.println(str);
    }

    @Test
    public void test4() {
        List<SubchainIpEntity> list = subchainIpService.selectList(new EntityWrapper<SubchainIpEntity>().eq("subchain_address", "0x5Fc8aE6BB3CfBb91d0617A58Bb115f991f245020"));
        Random random = new Random();
        int ran = random.nextInt(list.size());
        SubchainIpEntity subchainIpEntity = list.get(ran);
        System.out.println(subchainIpEntity.toString());
    }

    @Test
    public void test5() {
        subchainDao.insertSubList("123456");
        subchainDao.insertSubList("123789");
        subchainDao.insertSubList("321654");

        List<String> list = subchainDao.querySubList();
        subchainDao.deleteSubList("123456");
    }

    @Test
    public void test6() {
        List<UserEntity> userEntityList = userService.queryList(null);
        List<CurrencyEntity> currencyEntityList = currencyService.queryList(null);
        for (UserEntity userEntity : userEntityList) {
            for (CurrencyEntity currencyEntity : currencyEntityList) {
                UserAssetsEntity userAssetsEntity = userAssetsService.selectOne(
                        new EntityWrapper<UserAssetsEntity>()
                                .eq("currency_number", currencyEntity.getNumber())
                                .eq("user_id", userEntity.getUserId()));
                if (userAssetsEntity != null) {
                    continue;
                } else {
                    userAssetsEntity = new UserAssetsEntity();
                    userAssetsEntity.setBalance(new BigDecimal(0));
                    userAssetsEntity.setCurrencyNumber(currencyEntity.getNumber());
                    userAssetsEntity.setLock(new BigDecimal(0));
                    userAssetsEntity.setUserId(userEntity.getUserId());
                    userAssetsService.insert(userAssetsEntity);
                }
            }
        }
    }

    @Test
    public void test7() {
        List<UserEntity> userEntityList = userService.queryList(null);
        BigDecimal rate = new BigDecimal("1024");
        BigDecimal a1G = new BigDecimal("1").multiply(rate).multiply(rate).multiply(rate);
        for (UserEntity userEntity : userEntityList) {
            StoragePackageEntity storagePackageEntity = new StoragePackageEntity();
            storagePackageEntity.setStatus(-1);
            storagePackageEntity.setCreateDate(System.currentTimeMillis());
            storagePackageEntity.setCurrencyNumber("");
            storagePackageEntity.setExpireDate(-1l);
            storagePackageEntity.setFormatSize(Utils.formatFileSize(Long.valueOf(a1G.toEngineeringString())));
            storagePackageEntity.setPrice(new BigDecimal(0));
            storagePackageEntity.setRealSize(a1G);
            storagePackageEntity.setUseFormatSize("0");
            storagePackageEntity.setUserId(userEntity.getUserId());
            storagePackageEntity.setUserRealSize(new BigDecimal(0));
            storagePackageService.insert(storagePackageEntity);
        }

    }

    @Test
    public void test8() {
        StoragePackageEntity storagePackageEntity = storagePackageService.selectOne(
                new EntityWrapper<StoragePackageEntity>().eq("storage_package_id",3));
        System.out.println(storagePackageEntity);
    }
}
