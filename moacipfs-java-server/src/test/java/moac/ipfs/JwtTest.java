package moac.ipfs;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.baomidou.mybatisplus.mapper.EntityWrapper;
import io.swagger.annotations.ApiModel;
import moac.ipfs.common.httputils.OkHttpClients;
import moac.ipfs.common.httputils.OkHttpParam;
import moac.ipfs.common.httputils.OkhttpResult;
import moac.ipfs.common.utils.Constant;
import moac.ipfs.common.utils.IPUtils;
import moac.ipfs.common.utils.JwtUtils;
import moac.ipfs.common.utils.Utils;
import moac.ipfs.modules.back.user.entity.UserEntity;
import moac.ipfs.modules.back.user.service.UserService;
import moac.ipfs.modules.web.form.ImportAddressForm;
import moac.ipfs.modules.web.service.BaseWebService;
import org.apache.commons.lang.StringUtils;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.Map;


@RunWith(SpringRunner.class)
@SpringBootTest
public class JwtTest {
    @Autowired
    private JwtUtils jwtUtils;
    @Autowired
    private BaseWebService baseWebService;
    @Autowired
    private UserService userService;

    @Test
    public void test() {
        String token = jwtUtils.generateToken(1);
        System.out.println(token);
    }

    @Test
    public void test2() {
        OkHttpParam okHttpParam = new OkHttpParam();
        okHttpParam.setApiUrl(Constant.IMPORT_ADDRESS_URL);
        Map<String, String> map = new HashMap<>(16);
        String keystore = "{\"address\":\"7e3dcc26269b8fdd5a4ac5818a162a8f01403886\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"4a79393b4fc35d59d4567b256a9bc4b346ab0ba69ca97289aade2c44c8289e09\",\"cipherparams\":{\"iv\":\"c5275e52961cd3fe6408fd41d90d2ac5\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"db4e05ab51bf058634e8ef92876f6eda6282bf4948ac5d3f2f2a6fc3da3f7173\"},\"mac\":\"129d020e32e59fa88526d68f4f329642ba3e27bf8a79885244f37bfc4ea1b90c\"},\"id\":\"1bedc1a1-dcf4-4375-99d0-205b827c67cb\",\"version\":3}";
        map.put("importType", "KEYSTORE_TYPE");
        map.put("keystore", keystore);
        map.put("password", "123456");
        map.put("mnemonic", "");
        map.put("encryption", "");
        map.put("plaintextPrivateKey", "");
        OkhttpResult result = null;
        JSONObject json = null;
        try {
            result = OkHttpClients.post(okHttpParam, map, String.class, OkHttpClients.SYNCHRONIZE);
            json = JSON.parseObject(result.getResult());
            JSONObject jsonObject = json.getJSONObject("resultData");
            UserEntity userEntity = new UserEntity();
            userEntity.setCreateTime(System.currentTimeMillis());
            userEntity.setLastLoginIp("");
            userEntity.setPassword("");
            userEntity.setLastLoginTime(System.currentTimeMillis());
            userEntity.setUserAccount("IPFS" + Utils.getSerialNo());
            userEntity.setAddress(jsonObject.getString("address"));
            userEntity.setKeyStore(jsonObject.getString("keystore"));
            userService.insert(userEntity);
        } catch (
                Exception e) {
            e.printStackTrace();
        }
    }
}
