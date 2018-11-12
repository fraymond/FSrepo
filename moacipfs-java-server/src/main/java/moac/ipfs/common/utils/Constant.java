package moac.ipfs.common.utils;

/**
 * 常量
 * 
 * @author GZC
 * @email 57855143@qq.com
 * @date 2016年11月15日 下午1:23:52
 */
public class Constant {
	/** 超级管理员ID */
	public static final int SUPER_ADMIN = 1;
    /** 融云APPID */
    public static String RYAPPID = "8a216da863b5b9c20163c9c7c7f107f1";

	/**
	 * 菜单类型
	 * 
	 * @author GZC
	 * @email 57855143@qq.com
	 * @date 2016年11月15日 下午1:24:29
	 */
    public enum MenuType {
        /**
         * 目录
         */
    	CATALOG(0),
        /**
         * 菜单
         */
        MENU(1),
        /**
         * 按钮
         */
        BUTTON(2);

        private int value;

        MenuType(int value) {
            this.value = value;
        }

        public int getValue() {
            return value;
        }
    }
    

    /**
     * 云服务商
     */
    public enum CloudService {
        /**
         * 七牛云
         */
        QINIU(1),
        /**
         * 阿里云
         */
        ALIYUN(2),
        /**
         * 腾讯云
         */
        QCLOUD(3);

        private int value;

        CloudService(int value) {
            this.value = value;
        }

        public int getValue() {
            return value;
        }
    }

    /**
     * MOAC
     */
    public static final String MOAC = "MOAC";
    /**
     * keyStore导入地址方式
     */
    public static final String KEYSTORE_TYPE = "KEYSTORE_TYPE";
    /**
     * 助记词导入地址方式
     */
    public static final String MNEMONIC_TYPE = "MNEMONIC_TYPE";
    /**
     * 私钥导入地址方式
     */
    public static final String PLAINTEXTPRIVATEKEY_TYPE = "PLAINTEXTPRIVATEKEY_TYPE";

    /**
     * baseUrl
     */
    public static final String CREATE_base_URL = "http://47.107.99.26:8888/";
    /**
     *创建地址
     */
    public static final String CREATE_ADDRESS_URL = CREATE_base_URL + "createAddress";
    /**
     *导入地址
     */
    public static final String IMPORT_ADDRESS_URL = CREATE_base_URL + "importAddress";
    /**
     *查询文件
     */
    public static final String READFILE_URL = CREATE_base_URL + "readFile";
    /**
     *删除文件
     */
    public static final String REMOVEFILE_URL = CREATE_base_URL + "removeFile";
    /**
     *新增文件
     */
    public static final String SAVEFILE_URL = CREATE_base_URL + "saveFile";
    /**
     *删除存储子链
     */
    public static final String DELETEIPFS_URL = CREATE_base_URL + "deleteIpfs";
    /**
     *新增存储子链
     */
    public static final String ADDIPFS_URL = CREATE_base_URL + "addIpfs";
    /**
     *同步所有存储子链
     */
    public static final String GETALLIPFSINFO_URL = CREATE_base_URL + "getAllIpfsInfo";

}
