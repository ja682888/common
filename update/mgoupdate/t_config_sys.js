conn = new Mongo();
db = conn.getDB("test");


//1.数据库需要增加配置文件
db.t_config_sys.insert({
	id: 1, 
    newusercoin: 1000,
	newuserroomcard:20
});
