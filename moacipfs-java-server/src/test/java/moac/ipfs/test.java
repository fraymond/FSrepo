package moac.ipfs;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.fasterxml.jackson.core.JsonGenerationException;
import com.fasterxml.jackson.databind.JsonMappingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import moac.ipfs.common.utils.DateUtils;
import moac.ipfs.common.utils.MoacJsonRpc;
import moac.ipfs.common.utils.MoacUtils;
import moac.ipfs.common.utils.Utils;
import moac.ipfs.modules.web.form.ImportAddressForm;
import moac.ipfs.modules.web.service.impl.BaseWebServiceImpl;
import org.junit.Test;
import org.web3j.crypto.*;

import java.io.*;
import java.math.BigDecimal;
import java.math.BigInteger;
import java.math.RoundingMode;
import java.security.InvalidAlgorithmParameterException;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;
import java.time.ZoneOffset;
import java.time.ZonedDateTime;
import java.time.format.DateTimeFormatter;
import java.util.Scanner;

/**
 * @author: GZC
 * @create: 2018-11-20 17:56
 * @description:
 **/

public class test {

    /**
     * 创建moac地址
     */
    @Test
    public void test2() {
        try {
//            String fileName = this.getClass().getClassLoader().getResource("static").getPath();
//            System.out.println(fileName);
//            String file = WalletUtils.generateNewWalletFile(
//                    "123456",
//                    new File(fileName + "/"));
//            System.out.println(file);
//            delFile(file);

            ECKeyPair ecKeyPair = Keys.createEcKeyPair();
            WalletFile walletFile = Wallet.createStandard("123456", ecKeyPair);
//            String fileName = getWalletFileName(walletFile);
//            File destination = new File(destinationDirectory, fileName);
//            objectMapper.writeValue(destination, walletFile);
            JSONObject jsonObject = (JSONObject) JSON.toJSON(walletFile);
            System.out.println(jsonObject.toJSONString());
        } catch (CipherException e) {
            e.printStackTrace();
        } catch (InvalidAlgorithmParameterException e) {
            e.printStackTrace();
        } catch (NoSuchAlgorithmException e) {
            e.printStackTrace();
        } catch (NoSuchProviderException e) {
            e.printStackTrace();
        }
    }

    /**
     *
     */
    @Test
    public void test3() {
        try {
            String fileName = this.getClass().getClassLoader().getResource("static").getPath();
            System.out.println(fileName);
            Credentials credentials = WalletUtils.loadCredentials(
                    "123456",
                    fileName + "/" + "UTC--2018-11-22T08-10-55.854000000Z--ecf9dba25ea8cba96543841e079c8dd05b372b66.json");
            System.out.println(credentials);
            System.out.println(credentials.getEcKeyPair().getPrivateKey().toString(16));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    /**
     * keystore导入
     */
    @Test
    public void test4() {
        try {
            String fileUrl = createJsonFile("{\"address\":\"96ed14edb1d76121b92438c2c984ad58215cb30e\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"5fee0fc2ba17b5c2c5d3a43cf692128365182e68c88d7fbf97b7e475f91b77a8\",\"cipherparams\":{\"iv\":\"91595756d98645682df67c905e8c5c82\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"5445c993db6fb57ab89dc6b9fa9aad5491fc9558ca56952f8d2217928ddd0e58\"},\"mac\":\"1cc8c4050ec1069c94c74ea2f1c4ea4b78e17a143e353bdb452aba5f3cf3d8f3\"},\"id\":\"53a7ec26-78dd-454c-bbfa-b0334a41c645\",\"version\":3}");

            Credentials credentials = WalletUtils.loadCredentials(
                    "g123456",
                    fileUrl);
            delFile(fileUrl);
            System.out.println(credentials.getAddress());
            System.out.println(credentials.getEcKeyPair().getPrivateKey().toString(16));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    /**
     * 私钥导入
     */
    @Test
    public void Test5() {
        String privateKey = "e83c801a7b73b9741c8b3c26ac3349132ee89440a4d24fe5904f2c01b80cc6c5";

        ECKeyPair ecKeyPair = ECKeyPair.create(new BigInteger(privateKey, 16));
        try {
            WalletFile wallet = Wallet.createStandard("123456", ecKeyPair);
            ObjectMapper objectMapper = new ObjectMapper();
            DateTimeFormatter format = DateTimeFormatter.ofPattern("'UTC--'yyyy-MM-dd'T'HH-mm-ss.nVV'--'");
            ZonedDateTime now = ZonedDateTime.now(ZoneOffset.UTC);
            String fileName = now.format(format) + ".json";

            // 拼接文件完整路径
            String fileUrl = this.getClass().getClassLoader().getResource("static").getPath() + "/" + fileName;

            File destination = new File(fileUrl);
            objectMapper.writeValue(destination, wallet);
            System.out.println(fileUrl);
            delFile(fileUrl);
        } catch (CipherException e) {
            e.printStackTrace();
        } catch (JsonGenerationException e) {
            e.printStackTrace();
        } catch (JsonMappingException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }

    }


    /**
     * 生成.json格式文件
     */
    public String createJsonFile(String jsonString) {

        DateTimeFormatter format = DateTimeFormatter.ofPattern("'UTC--'yyyy-MM-dd'T'HH-mm-ss.nVV'--'");
        ZonedDateTime now = ZonedDateTime.now(ZoneOffset.UTC);
        String fileName = now.format(format) + ".json";

        // 拼接文件完整路径
        String fileUrl = this.getClass().getClassLoader().getResource("static").getPath() + "/" + fileName;

        // 生成json格式文件
        try {
            // 保证创建一个新文件
            File file = new File(fileUrl);
            if (!file.getParentFile().exists()) { // 如果父目录不存在，创建父目录
                file.getParentFile().mkdirs();
            }
            if (file.exists()) { // 如果已存在,删除旧文件
                file.delete();
            }
            file.createNewFile();


            // 将格式化后的字符串写入文件
            Writer write = new OutputStreamWriter(new FileOutputStream(file), "UTF-8");
            write.write(jsonString);
            write.flush();
            write.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
        return fileUrl;
    }

    public void delFile(String fileUrl) {
        File file = new File(fileUrl);
        if (file.exists() && file.isFile())
            file.delete();
    }

    @Test
    public void queryMainChainBalance() {
        String address = "0xd58592114ebd97525856929d5c662b72d58b767b";
        String result =  MoacJsonRpc.queryMainChainBalance(address);
        String balance = result.substring(2);
        BigDecimal number = MoacUtils.getBaclanceNumber(balance,18);
        System.out.println(number.setScale(4,BigDecimal.ROUND_FLOOR).toPlainString());
    }

    @Test
    public void querySubChainBalance() {
        String address = "0xd58592114ebd97525856929d5c662b72d58b767b";
        String result =  MoacJsonRpc.querySubChainBalance(address);
        BigDecimal number = MoacUtils.getSubBaclanceNumber(result,18);
        System.out.println(number.setScale(4,BigDecimal.ROUND_FLOOR).toPlainString());
    }

    @Test
    public void queryMainChainBlock() {
        String result =  MoacJsonRpc.queryMainChainBlock();
        BigInteger big = new BigInteger(result.substring(2),16);
        System.out.println(big.toString(10));
    }

    @Test
    public void querySubChainBlock() {
        String result =  MoacJsonRpc.querySubChainBlock();
        System.out.println(result);
    }

    @Test
    public void tesdt() {
        long result = Utils.formatFileSize("500G");
        System.out.println(result);
        String str = Utils.formatFileSize(result);
        System.out.println(str);
    }

    @Test
    public void tesdt2() {
        String str= "0.5";
        Long l = Long.valueOf(str);
        System.out.println(l);
    }
}
