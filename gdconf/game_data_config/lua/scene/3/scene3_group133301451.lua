-- 基础信息
local base_info = {
	group_id = 133301451
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 451001, monster_id = 20011001, pos = { x = -498.424, y = 164.700, z = 3298.004 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 33, drop_tag = "史莱姆", area_id = 22 },
	{ config_id = 451002, monster_id = 20011001, pos = { x = -485.607, y = 164.700, z = 3292.065 }, rot = { x = 0.000, y = 131.915, z = 0.000 }, level = 33, drop_tag = "史莱姆", area_id = 22 },
	{ config_id = 451003, monster_id = 20011001, pos = { x = -480.618, y = 164.700, z = 3308.111 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 33, drop_tag = "史莱姆", area_id = 22 },
	{ config_id = 451004, monster_id = 20011101, pos = { x = -486.385, y = 164.700, z = 3295.140 }, rot = { x = 0.000, y = 124.636, z = 0.000 }, level = 33, drop_tag = "大史莱姆", area_id = 22 }
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
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
		monsters = { 451001, 451002, 451003, 451004 },
		gadgets = { },
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