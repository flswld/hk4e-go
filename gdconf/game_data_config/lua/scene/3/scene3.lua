-- 地图配置
scene_config = {
	begin_pos = { x = -8192.0, z = -6144.0 },
	size = { x = 11264.0, z = 12288.0 },
	born_pos = { x = 2747.562, y = 194.633, z = -1719.386 },
	born_rot = { x = 0.000, y = 307.000, z = 0.000 },
	die_y = -280,
	vision_anchor = { x = -8192.0, z = -8192.0 }
}

-- 所有的区块
blocks = { 3001, 3002, 3003, 3004, 3005, 3006, 3007, 3008, 3101, 3102, 3103, 3104, 3105, 3106, 3107, 3108, 3109, 3110, 3113, 3204, 3205, 3207, 3208, 3209, 3210, 3211, 3212, 3213, 3214, 3215, 3216, 3217, 3218, 3219, 3220, 3221, 3222, 3223, 3224, 3225, 3226, 3301, 3302, 3303, 3304, 3305, 3306, 3307, 3308, 3309, 3310, 3311, 3312, 3313, 3314, 3315 }

-- 所有的区块范围坐标
block_rects = {
	{ min = { x = 1024.0, z = -2048.0 }, max = { x = 2048.0, z = -1024.0 } },
	{ min = { x = 1024.0, z = -1024.0 }, max = { x = 2048.0, z = 0.0 } },
	{ min = { x = 2048.0, z = -2048.0 }, max = { x = 3072.0, z = -1024.0 } },
	{ min = { x = 2048.0, z = -1024.0 }, max = { x = 3072.0, z = 0.0 } },
	{ min = { x = 1024.0, z = -3072.0 }, max = { x = 2048.0, z = -2048.0 } },
	{ min = { x = 2048.0, z = -3072.0 }, max = { x = 3072.0, z = -2048.0 } },
	{ min = { x = 2048.0, z = 0.0 }, max = { x = 3072.0, z = 1024.0 } },
	{ min = { x = 0.0, z = -2048.0 }, max = { x = 1024.0, z = -1024.0 } },
	{ min = { x = 1024.0, z = 1024.0 }, max = { x = 2048.0, z = 2048.0 } },
	{ min = { x = 1024.0, z = 0.0 }, max = { x = 2048.0, z = 1024.0 } },
	{ min = { x = 0.0, z = 1024.0 }, max = { x = 1024.0, z = 2048.0 } },
	{ min = { x = 0.0, z = 0.0 }, max = { x = 1024.0, z = 1024.0 } },
	{ min = { x = 0.0, z = -1024.0 }, max = { x = 1024.0, z = 0.0 } },
	{ min = { x = -1024.0, z = 1024.0 }, max = { x = 0.0, z = 2048.0 } },
	{ min = { x = -1024.0, z = 0.0 }, max = { x = 0.0, z = 1024.0 } },
	{ min = { x = -1024.0, z = -1024.0 }, max = { x = 0.0, z = 0.0 } },
	{ min = { x = -2048.0, z = 1024.0 }, max = { x = -1024.0, z = 2048.0 } },
	{ min = { x = -2048.0, z = 0.0 }, max = { x = -1024.0, z = 1024.0 } },
	{ min = { x = -1024.0, z = -2048.0 }, max = { x = 0.0, z = -1024.0 } },
	{ min = { x = -2048.0, z = -3072.0 }, max = { x = -1024.0, z = -2048.0 } },
	{ min = { x = -2048.0, z = -4096.0 }, max = { x = -1024.0, z = -3072.0 } },
	{ min = { x = -3072.0, z = -2048.0 }, max = { x = -2048.0, z = -1024.0 } },
	{ min = { x = -3072.0, z = -3072.0 }, max = { x = -2048.0, z = -2048.0 } },
	{ min = { x = -3072.0, z = -4096.0 }, max = { x = -2048.0, z = -3072.0 } },
	{ min = { x = -4096.0, z = -1024.0 }, max = { x = -3072.0, z = 0.0 } },
	{ min = { x = -4096.0, z = -2048.0 }, max = { x = -3072.0, z = -1024.0 } },
	{ min = { x = -4096.0, z = -3072.0 }, max = { x = -3072.0, z = -2048.0 } },
	{ min = { x = -4096.0, z = -4096.0 }, max = { x = -3072.0, z = -3072.0 } },
	{ min = { x = -5120.0, z = -1024.0 }, max = { x = -4096.0, z = 0.0 } },
	{ min = { x = -5120.0, z = -2048.0 }, max = { x = -4096.0, z = -1024.0 } },
	{ min = { x = -5120.0, z = -3072.0 }, max = { x = -4096.0, z = -2048.0 } },
	{ min = { x = -5120.0, z = -4096.0 }, max = { x = -4096.0, z = -3072.0 } },
	{ min = { x = -1024.0, z = -5120.0 }, max = { x = 0.0, z = -4096.0 } },
	{ min = { x = -2048.0, z = -5120.0 }, max = { x = -1024.0, z = -4096.0 } },
	{ min = { x = -3072.0, z = -5120.0 }, max = { x = -2048.0, z = -4096.0 } },
	{ min = { x = -4096.0, z = -5120.0 }, max = { x = -3072.0, z = -4096.0 } },
	{ min = { x = -5120.0, z = -5120.0 }, max = { x = -4096.0, z = -4096.0 } },
	{ min = { x = -6144.0, z = -3072.0 }, max = { x = -5120.0, z = -2048.0 } },
	{ min = { x = -6144.0, z = -4096.0 }, max = { x = -5120.0, z = -3072.0 } },
	{ min = { x = -7168.0, z = -3072.0 }, max = { x = -6144.0, z = -2048.0 } },
	{ min = { x = -7168.0, z = -4096.0 }, max = { x = -6144.0, z = -3072.0 } },
	{ min = { x = -1024.0, z = 3072.0 }, max = { x = 0.0, z = 4096.0 } },
	{ min = { x = -1024.0, z = 2048.0 }, max = { x = 0.0, z = 3072.0 } },
	{ min = { x = -2048.0, z = 3072.0 }, max = { x = -1024.0, z = 4096.0 } },
	{ min = { x = -2048.0, z = 2048.0 }, max = { x = -1024.0, z = 3072.0 } },
	{ min = { x = -3072.0, z = 3072.0 }, max = { x = -2048.0, z = 4096.0 } },
	{ min = { x = -3072.0, z = 2048.0 }, max = { x = -2048.0, z = 3072.0 } },
	{ min = { x = -2048.0, z = 5120.0 }, max = { x = -1024.0, z = 6144.0 } },
	{ min = { x = -2048.0, z = 4096.0 }, max = { x = -1024.0, z = 5120.0 } },
	{ min = { x = -3072.0, z = 5120.0 }, max = { x = -2048.0, z = 6144.0 } },
	{ min = { x = -3072.0, z = 4096.0 }, max = { x = -2048.0, z = 5120.0 } },
	{ min = { x = -4096.0, z = 5120.0 }, max = { x = -3072.0, z = 6144.0 } },
	{ min = { x = -4096.0, z = 4096.0 }, max = { x = -3072.0, z = 5120.0 } },
	{ min = { x = -1024.0, z = 5120.0 }, max = { x = 0.0, z = 6144.0 } },
	{ min = { x = -1024.0, z = 4096.0 }, max = { x = 0.0, z = 5120.0 } },
	{ min = { x = 0.0, z = 2048.0 }, max = { x = 1024.0, z = 3072.0 } }
}

-- Dummy Points
dummy_points = { "dummy_points" }

-- Routes
routes_config = { "routes" }