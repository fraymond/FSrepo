package moac.ipfs;

import com.alibaba.fastjson.JSON;
import moac.ipfs.common.utils.BeanUtils;
import moac.ipfs.common.utils.RedisUtils;
import moac.ipfs.modules.back.sys.dao.SysUserDao;
import moac.ipfs.modules.back.sys.entity.SysUserEntity;
import org.apache.commons.lang.builder.ToStringBuilder;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

@RunWith(SpringRunner.class)
@SpringBootTest
public class RedisTest {
	@Autowired
	private RedisUtils redisUtils;
	@Autowired
	private SysUserDao sysUserDao;
	@Autowired
	private BeanUtils beanUtils;

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
		redisUtils.set("list",str,RedisUtils.FIVEMINUTE_EXPIRE);
		String s = redisUtils.get("list");
		List<String> stringList = JSON.parseArray(s,String.class);
		System.out.println(stringList.toString());



	}
}
