-- 基础信息
local base_info = {
	group_id = 133305032
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 32001, monster_id = 25010301, pos = { x = -2172.918, y = 318.650, z = 4027.044 }, rot = { x = 0.000, y = 144.949, z = 0.000 }, level = 32, drop_tag = "盗宝团", area_id = 26 },
	{ config_id = 32002, monster_id = 25020201, pos = { x = -2178.950, y = 322.382, z = 4025.442 }, rot = { x = 0.000, y = 123.469, z = 0.000 }, level = 32, drop_tag = "盗宝团", area_id = 26 },
	{ config_id = 32003, monster_id = 25070101, pos = { x = -2171.062, y = 320.090, z = 4013.261 }, rot = { x = 0.000, y = 83.839, z = 0.000 }, level = 32, drop_tag = "盗宝团", area_id = 26 }
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 32006, gadget_id = 70220050, pos = { x = -2173.740, y = 319.937, z = 4015.852 }, rot = { x = 6.236, y = 359.699, z = 356.407 }, level = 32, area_id = 26 },
	{ config_id = 32007, gadget_id = 70220052, pos = { x = -2173.069, y = 319.841, z = 4016.323 }, rot = { x = 8.105, y = 359.693, z = 359.024 }, level = 32, area_id = 26 },
	{ config_id = 32008, gadget_id = 70330397, pos = { x = -2170.876, y = 320.798, z = 4006.999 }, rot = { x = 4.926, y = 339.642, z = 0.392 }, level = 32, area_id = 26 }
}

-- 区域
regions = {
}

-- 触发器
triggers = {
}

-- 变量
variables = {
}

--================================================================
-- 
-- 初始化配置
-- 
--================================================================

-- 初始化时创建
init_config = {
	suite = 1,
	end_suite = 0,
	rand_suite = false
}

--================================================================
-- 
-- 小组配置
-- 
--================================================================

suites = {
	{
		-- suite_id = 1,
		-- description = ,
		monsters = { 32001, 32002, 32003 },
		gadgets = { 32006, 32007, 32008 },
		regions = { },
		triggers = { },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================