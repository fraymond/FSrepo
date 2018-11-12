package moac.ipfs.modules.web.service;


import com.baomidou.mybatisplus.service.IService;
import moac.ipfs.modules.back.user.entity.UserEntity;
import moac.ipfs.modules.back.user.entity.UserFileEntity;
import moac.ipfs.modules.web.form.ImportAddressForm;
import moac.ipfs.modules.web.form.CreateAddressForm;
import moac.ipfs.modules.web.form.FileParamsForm;
import moac.ipfs.modules.web.form.QueryFileListForm;

import javax.servlet.http.HttpServletRequest;
import java.util.List;

/**
 * 用户
 * 
 * @author GZC
 * @email 57855143@qq.com
 * @date 2017-03-23 15:22:06
 */
public interface BaseWebService extends IService<UserEntity> {

	/**
	 * 导入钱包
	 * @param importAddressForm
	 * @param request
	 * @return
	 */
	 UserEntity importAddressService(ImportAddressForm importAddressForm,HttpServletRequest request);


	/**
	 * 创建地址
	 * @param form
	 * @param request
	 * @return
	 */
	UserEntity createAddressService(CreateAddressForm form, HttpServletRequest request);


	/**
	 * 查看文件
	 * @param form
	 * @return
	 */
	String queryFileByHashService(FileParamsForm form);

	/**
	 * 删除文件
	 * @param form
	 * @return
	 */
	void deleteFileService(FileParamsForm form);

	/**
	 * 新增文件
	 * @param form
	 */
	void addFileService(FileParamsForm form);
}
