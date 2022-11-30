-- 基础信息
local base_info = {
	group_id = 250032034
}

-- Trigger变量
local defs = {
	group_id = 250032034,
	gear_group_id = 250032037
}

-- DEFS_MISCS
tide_defs = {
	[1] = { {sum = 20, min = 7, max = 7,suite=2},{sum = 20, min = 7, max = 7,suite=3},},
	[2] = { {sum = 20, min = 7, max = 7,suite=4},{sum = 20, min = 7, max = 7,suite=5},},
}

package_tide_defs={
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	[34001] = { config_id = 34001, monster_id = 20011302, pos = { x = -19.059, y = -0.350, z = -53.124 }, rot = { x = 0.000, y = 172.985, z = 0.000 }, level = 1 },
	[34002] = { config_id = 34002, monster_id = 20011302, pos = { x = -16.901, y = -0.350, z = -53.166 }, rot = { x = 0.000, y = 172.985, z = 0.000 }, level = 1 },
	[34003] = { config_id = 34003, monster_id = 20011203, pos = { x = -17.935, y = -0.347, z = -50.945 }, rot = { x = 0.000, y = 172.985, z = 0.000 }, level = 1 },
	[34004] = { config_id = 34004, monster_id = 20011203, pos = { x = -16.605, y = -0.349, z = -51.894 }, rot = { x = 0.000, y = 172.985, z = 0.000 }, level = 1 },
	[34005] = { config_id = 34005, monster_id = 20011203, pos = { x = -20.613, y = -0.356, z = -51.131 }, rot = { x = 0.000, y = 172.985, z = 0.000 }, level = 1 },
	[34006] = { config_id = 34006, monster_id = 20010903, pos = { x = -16.678, y = -0.380, z = -82.704 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	[34007] = { config_id = 34007, monster_id = 20010903, pos = { x = -18.446, y = -0.350, z = -85.176 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	[34008] = { config_id = 34008, monster_id = 20010803, pos = { x = -17.601, y = -0.350, z = -84.008 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	[34009] = { config_id = 34009, monster_id = 20010803, pos = { x = -15.348, y = -0.349, z = -85.233 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	[34010] = { config_id = 34010, monster_id = 20010803, pos = { x = -16.120, y = -0.350, z = -84.127 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	[34011] = { config_id = 34011, monster_id = 20011102, pos = { x = -19.059, y = -0.350, z = -53.124 }, rot = { x = 0.000, y = 172.985, z = 0.000 }, level = 1 },
	[34012] = { config_id = 34012, monster_id = 20011102, pos = { x = -16.901, y = -0.350, z = -53.166 }, rot = { x = 0.000, y = 172.985, z = 0.000 }, level = 1 },
	[34013] = { config_id = 34013, monster_id = 20011002, pos = { x = -17.935, y = -0.347, z = -50.945 }, rot = { x = 0.000, y = 172.985, z = 0.000 }, level = 1 },
	[34014] = { config_id = 34014, monster_id = 20011002, pos = { x = -16.605, y = -0.349, z = -51.894 }, rot = { x = 0.000, y = 172.985, z = 0.000 }, level = 1 },
	[34015] = { config_id = 34015, monster_id = 20011002, pos = { x = -20.613, y = -0.356, z = -51.131 }, rot = { x = 0.000, y = 172.985, z = 0.000 }, level = 1 },
	[34016] = { config_id = 34016, monster_id = 20010602, pos = { x = -16.120, y = -0.350, z = -84.127 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	[34017] = { config_id = 34017, monster_id = 20010703, pos = { x = -15.348, y = -0.349, z = -85.233 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	[34018] = { config_id = 34018, monster_id = 20010502, pos = { x = -17.601, y = -0.350, z = -84.008 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	[34019] = { config_id = 34019, monster_id = 20010502, pos = { x = -18.446, y = -0.350, z = -85.176 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	[34020] = { config_id = 34020, monster_id = 20010502, pos = { x = -16.678, y = -0.380, z = -82.704 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	[34021] = { config_id = 34021, monster_id = 21010102, pos = { x = -16.388, y = -0.356, z = -51.097 }, rot = { x = 0.000, y = 172.985, z = 0.000 }, level = 1 },
	[34022] = { config_id = 34022, monster_id = 21010102, pos = { x = -20.613, y = -0.356, z = -51.131 }, rot = { x = 0.000, y = 172.985, z = 0.000 }, level = 1 },
	[34023] = { config_id = 34023, monster_id = 21010102, pos = { x = -18.815, y = -0.356, z = -86.092 }, rot = { x = 0.000, y = 354.121, z = 0.000 }, level = 1 },
	[34024] = { config_id = 34024, monster_id = 21010102, pos = { x = -14.591, y = -0.356, z = -86.141 }, rot = { x = 0.000, y = 354.121, z = 0.000 }, level = 1 },
	[34025] = { config_id = 34025, monster_id = 21010102, pos = { x = -16.388, y = -0.356, z = -51.097 }, rot = { x = 0.000, y = 172.985, z = 0.000 }, level = 1 },
	[34026] = { config_id = 34026, monster_id = 21010102, pos = { x = -20.613, y = -0.356, z = -51.131 }, rot = { x = 0.000, y = 172.985, z = 0.000 }, level = 1 },
	[34027] = { config_id = 34027, monster_id = 21010102, pos = { x = -18.815, y = -0.356, z = -86.092 }, rot = { x = 0.000, y = 354.121, z = 0.000 }, level = 1 },
	[34028] = { config_id = 34028, monster_id = 21010102, pos = { x = -14.591, y = -0.356, z = -86.141 }, rot = { x = 0.000, y = 354.121, z = 0.000 }, level = 1 }
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

-- 点位
points = {
	{ config_id = 34029, pos = { x = -18.714, y = -0.352, z = -54.720 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 2 },
	{ config_id = 34030, pos = { x = -17.271, y = -0.355, z = -54.667 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 2 },
	{ config_id = 34031, pos = { x = -15.653, y = -0.359, z = -54.696 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 2 },
	{ config_id = 34032, pos = { x = -19.025, y = -0.350, z = -53.068 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 2 },
	{ config_id = 34033, pos = { x = -17.347, y = -0.348, z = -52.974 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 2 },
	{ config_id = 34034, pos = { x = -15.481, y = -0.355, z = -52.792 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 2 },
	{ config_id = 34035, pos = { x = -17.463, y = -0.347, z = -51.052 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 2 },
	{ config_id = 34036, pos = { x = -14.721, y = -0.347, z = -82.318 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, tag = 4 },
	{ config_id = 34037, pos = { x = -16.384, y = -0.370, z = -82.404 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, tag = 4 },
	{ config_id = 34038, pos = { x = -18.056, y = -0.350, z = -82.372 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, tag = 4 },
	{ config_id = 34039, pos = { x = -14.623, y = -0.348, z = -84.302 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, tag = 4 },
	{ config_id = 34040, pos = { x = -16.646, y = -0.350, z = -84.411 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, tag = 4 },
	{ config_id = 34041, pos = { x = -18.596, y = -0.350, z = -84.504 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, tag = 4 },
	{ config_id = 34042, pos = { x = -14.697, y = -0.355, z = -86.663 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, tag = 4 },
	{ config_id = 34043, pos = { x = -16.957, y = -0.349, z = -86.861 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, tag = 4 },
	-- StartPoint1
	{ config_id = 34044, pos = { x = -13.740, y = -0.347, z = -52.769 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 2 },
	-- StartPoint1
	{ config_id = 34045, pos = { x = -13.740, y = -0.347, z = -54.337 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 2 }
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
		gadgets = { },
		regions = { },
		triggers = { },
		rand_weight = 100
	},
	{
		-- suite_id = 2,
		-- description = Option1-1-C,
		monsters = { 34001, 34002, 34003, 34004, 34005, 34021, 34022 },
		gadgets = { },
		regions = { },
		triggers = { },
		rand_weight = 100
	},
	{
		-- suite_id = 3,
		-- description = Option1-2-C,
		monsters = { 34006, 34007, 34008, 34009, 34010, 34023, 34024 },
		gadgets = { },
		regions = { },
		triggers = { },
		rand_weight = 100
	},
	{
		-- suite_id = 4,
		-- description = Option2-1-D,
		monsters = { 34011, 34012, 34013, 34014, 34015, 34025, 34026 },
		gadgets = { },
		regions = { },
		triggers = { },
		rand_weight = 100
	},
	{
		-- suite_id = 5,
		-- description = Option2-2-D,
		monsters = { 34016, 34017, 34018, 34019, 34020, 34027, 34028 },
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

require "TowerDefense_Monster02"