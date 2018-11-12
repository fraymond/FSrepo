package moac.ipfs;


import com.alibaba.fastjson.JSON;
import moac.ipfs.common.utils.RedisUtils;
import moac.ipfs.modules.back.subchain.entity.SubchainEntity;
import moac.ipfs.modules.back.subchain.service.SubchainService;
import moac.ipfs.modules.back.sys.entity.SysUserEntity;
import moac.ipfs.modules.web.service.impl.BaseWebServiceImpl;
import moac.ipfs.service.DataSourceTestService;
import org.apache.commons.lang.builder.ToStringBuilder;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


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

    @Test
    public void test(){
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
    public void test2(){
        Map<String, Object> params1 = new HashMap<>(16);
        params1.put("fileSize", 10000);
        SubchainEntity subchainEntity = subchainService.queryAddFileList(params1);
        List<String> recentlyUsedList = new ArrayList<>();
        recentlyUsedList.add(subchainEntity.getSubchainAddress());
        String str = JSON.toJSONString(recentlyUsedList);
        redisUtils.set("recentlyUsed", str, -1);
        List<String> recentlyUsed = JSON.parseArray(redisUtils.get("recentlyUsed"), String.class);
        Map<String, Object> params = new HashMap<>(16);
        if (recentlyUsed != null){
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

}
