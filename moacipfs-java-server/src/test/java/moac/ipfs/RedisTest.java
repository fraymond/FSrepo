package moac.ipfs;

import com.alibaba.fastjson.JSON;
import moac.ipfs.common.utils.BeanUtils;
import moac.ipfs.common.utils.RedisUtils;
import moac.ipfs.modules.back.sys.dao.SysUserDao;
import moac.ipfs.modules.back.sys.entity.SysUserEntity;
import moac.ipfs.modules.web.form.StoragePackageParamsFrom;
import moac.ipfs.modules.web.service.BaseWebService;
import org.apache.commons.lang.builder.ToStringBuilder;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import java.math.BigDecimal;
import java.math.BigInteger;
import java.math.RoundingMode;
import java.util.*;

@RunWith(SpringRunner.class)
@SpringBootTest
public class RedisTest {
    @Autowired
    private RedisUtils redisUtils;
    @Autowired
    private SysUserDao sysUserDao;
    @Autowired
    private BeanUtils beanUtils;
    @Autowired
    private BaseWebService baseWebService;

    /**
     * keystore导入
     */
    @Test
    public void test11() {
        try {
            String keyStore = "{\"address\":\"d58592114ebd97525856929d5c662b72d58b767b\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"db96d030406419a3ca0d6e6901b3b688ad4c6f34376f048c8ed56f39d1b37169\",\"cipherparams\":{\"iv\":\"9e6aee2025bd919866da952d3df40566\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"f22f59cac654619c29efba26fef1c1b894691b0e2e6a2bf903104dae96d824c4\"},\"mac\":\"af807312730a542925badcb3830853a43324c8b19b27cbefff7e72f150e6576f\"},\"id\":\"d4bfe897-6bcb-4f5d-9166-cc9a4b5ba488\",\"version\":3}";
            StoragePackageParamsFrom storagePackageEntity = new StoragePackageParamsFrom();
            baseWebService.subChainTransService(storagePackageEntity);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    @Test
    public void contextLoads() {
        SysUserEntity user = new SysUserEntity();
        user.setEmail("qqq@qq.com");
        redisUtils.set("user", user);

        System.out.println(ToStringBuilder.reflectionToString(redisUtils.get("user", SysUserEntity.class)));
    }

    @Test
    public void contextLoads2() {
        SysUserEntity user = sysUserDao.selectById(1);
        System.out.println(user.toString());
        System.out.println(ToStringBuilder.reflectionToString(redisUtils.get("user", SysUserEntity.class)));

    }

    @Test
    public void contextLoads3() {
        SysUserEntity user = sysUserDao.selectById(1);
        System.out.println(user);
        try {
            SysUserEntity userEntity = beanUtils.deepClone(user);
            System.out.println(userEntity);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    @Test
    public void contextLoads4() {
        List<String> list = new ArrayList<>(16);
        list.add("asd");
        list.add("s");
        list.add("223");
        list.add("dadadadada");
        list.add("ssssss");
        String str = JSON.toJSONString(list);
        redisUtils.set("list", str, RedisUtils.FIVEMINUTE_EXPIRE);
        String s = redisUtils.get("list");
        List<String> stringList = JSON.parseArray(s, String.class);
        System.out.println(stringList.toString());
    }

    @Test
    public void contextLoads5() {
        String strPart = "10000000000000000000000";
        System.out.println(new BigInteger(strPart, 10).toString(32));
    }

    @Test
    public void contextLoads6() {
        BigDecimal total = new BigDecimal("200000000");
        BigDecimal month = total.divide(new BigDecimal(6), 4, RoundingMode.FLOOR);
        double filestorm1;
        double filestorm2;
        double filestorm3;
        double filestorm4;
        double filestorm5;
        double filestorm6;
        double filestorm7;

        Scanner in = new Scanner(System.in);
        System.out.println("输入filestorm1：");
        filestorm1 = in.nextInt();
        System.out.println("输入filestorm2：");
        filestorm2 = in.nextInt();
        System.out.println("输入filestorm3：");
        filestorm3 = in.nextInt();
        System.out.println("输入filestorm4：");
        filestorm4 = in.nextInt();
        System.out.println("输入filestorm5：");
        filestorm5 = in.nextInt();
        System.out.println("输入filestorm6：");
        filestorm6 = in.nextInt();
        System.out.println("输入filestorm7：");
        filestorm7 = in.nextInt();

        double value = month.doubleValue() /
                (filestorm1 * 1 + filestorm2 * 2 + filestorm3 * 4 + filestorm4 * 8 + filestorm5 * 16 + filestorm6 * 32 + filestorm7 * 64);

    }

}
