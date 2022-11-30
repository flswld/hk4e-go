-- 基础信息
local base_info = {
	group_id = 133307055
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 55001, gadget_id = 70330254, pos = { x = -1852.396, y = 168.272, z = 5800.508 }, rot = { x = 13.639, y = 236.300, z = 346.043 }, level = 32, area_id = 27 },
	{ config_id = 55002, gadget_id = 70330254, pos = { x = -1848.831, y = 172.167, z = 5740.882 }, rot = { x = 16.326, y = 276.070, z = 354.815 }, level = 32, area_id = 27 },
	{ config_id = 55003, gadget_id = 70330254, pos = { x = -1912.417, y = 161.017, z = 5703.584 }, rot = { x = 0.000, y = 305.476, z = 0.000 }, level = 32, area_id = 27 },
	{ config_id = 55004, gadget_id = 70330254, pos = { x = -1908.238, y = 162.524, z = 5707.378 }, rot = { x = 17.174, y = 307.617, z = 10.690 }, level = 32, area_id = 27 },
	{ config_id = 55005, gadget_id = 70330254, pos = { x = -1859.875, y = 165.610, z = 5803.664 }, rot = { x = 11.177, y = 236.917, z = 346.718 }, level = 32, area_id = 27 },
	{ config_id = 55006, gadget_id = 70330254, pos = { x = -1801.298, y = 183.179, z = 5789.599 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, area_id = 27 },
	{ config_id = 55007, gadget_id = 70330254, pos = { x = -1801.034, y = 183.276, z = 5782.432 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, area_id = 27 },
	{ config_id = 55008, gadget_id = 70330254, pos = { x = -1760.003, y = 191.987, z = 5828.343 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 55009, gadget_id = 70330254, pos = { x = -1833.445, y = 164.169, z = 5881.029 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, area_id = 27 },
	{ config_id = 55010, gadget_id = 70330254, pos = { x = -1827.347, y = 166.060, z = 5878.610 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, area_id = 27 },
	{ config_id = 55011, gadget_id = 70330254, pos = { x = -1879.030, y = 166.935, z = 5755.836 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, area_id = 27 },
	{ config_id = 55012, gadget_id = 70330254, pos = { x = -1854.186, y = 172.953, z = 5716.687 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, area_id = 27 },
	{ config_id = 55013, gadget_id = 70330254, pos = { x = -1913.470, y = 155.822, z = 5754.537 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, area_id = 27 },
	{ config_id = 55014, gadget_id = 70330254, pos = { x = -1904.109, y = 154.430, z = 5797.746 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, area_id = 27 }
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
		monsters = { },
		gadgets = { 55001, 55002, 55003, 55004, 55005, 55006, 55007, 55008, 55009, 55010, 55011, 55012, 55013, 55014 },
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