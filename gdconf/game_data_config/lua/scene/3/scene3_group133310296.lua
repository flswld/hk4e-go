-- 基础信息
local base_info = {
	group_id = 133310296
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
	{ config_id = 296001, gadget_id = 70330432, pos = { x = -2372.706, y = 104.808, z = 4872.892 }, rot = { x = 327.727, y = 43.308, z = 0.000 }, level = 32, area_id = 27 },
	{ config_id = 296002, gadget_id = 70330432, pos = { x = -2337.121, y = 103.822, z = 4875.341 }, rot = { x = 41.729, y = 339.244, z = 0.000 }, level = 32, area_id = 27 },
	{ config_id = 296003, gadget_id = 70330432, pos = { x = -2334.265, y = 104.378, z = 4915.639 }, rot = { x = 322.236, y = 215.482, z = 0.000 }, level = 32, area_id = 27 },
	{ config_id = 296004, gadget_id = 70330432, pos = { x = -2363.735, y = 103.891, z = 4897.840 }, rot = { x = 327.870, y = 160.564, z = 19.096 }, level = 32, area_id = 27 },
	{ config_id = 296005, gadget_id = 70330432, pos = { x = -2359.140, y = 103.935, z = 4854.688 }, rot = { x = 346.440, y = 178.948, z = 294.257 }, level = 32, area_id = 27 },
	{ config_id = 296006, gadget_id = 70330432, pos = { x = -2347.713, y = 104.684, z = 4844.245 }, rot = { x = 32.059, y = 0.000, z = 0.000 }, level = 32, area_id = 26 },
	{ config_id = 296007, gadget_id = 70330432, pos = { x = -2342.424, y = 103.694, z = 4892.207 }, rot = { x = 4.011, y = 190.981, z = 318.015 }, level = 32, area_id = 27 }
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
		gadgets = { 296001, 296002, 296003, 296004, 296005, 296006, 296007 },
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