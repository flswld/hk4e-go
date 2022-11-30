-- 基础信息
local base_info = {
	group_id = 133221070
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
	{ config_id = 70005, npc_id = 12490, pos = { x = -3027.591, y = 236.635, z = -4445.708 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, area_id = 11 },
	{ config_id = 70007, npc_id = 12490, pos = { x = -3027.591, y = 236.635, z = -4445.708 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, area_id = 11 }
}

-- 装置
gadgets = {
	{ config_id = 70001, gadget_id = 70710443, pos = { x = -3027.591, y = 236.010, z = -4445.708 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1, area_id = 11 },
	{ config_id = 70002, gadget_id = 70710645, pos = { x = -3030.373, y = 236.010, z = -4445.537 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, level = 1, area_id = 11 },
	{ config_id = 70003, gadget_id = 70710645, pos = { x = -3030.373, y = 236.010, z = -4445.537 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, level = 1, area_id = 11 },
	{ config_id = 70004, gadget_id = 70710490, pos = { x = -3030.421, y = 236.010, z = -4446.801 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1, area_id = 11 },
	{ config_id = 70006, gadget_id = 70710490, pos = { x = -3030.421, y = 236.010, z = -4446.801 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1, area_id = 11 },
	{ config_id = 70008, gadget_id = 70710443, pos = { x = -3027.591, y = 236.010, z = -4445.708 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1, area_id = 11 }
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
		gadgets = { 70001, 70002, 70004 },
		regions = { },
		triggers = { },
		npcs = { 70005 },
		rand_weight = 100
	},
	{
		-- suite_id = 2,
		-- description = ,
		monsters = { },
		gadgets = { 70003, 70006, 70008 },
		regions = { },
		triggers = { },
		npcs = { 70007 },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================