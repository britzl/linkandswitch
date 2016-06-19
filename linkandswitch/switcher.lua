local block = require "linkandswitch.block"


local M = {}

local function mark_match(match)
	if #match >= 3 then
		for _,v in pairs(match) do
			v.match = true
		end
	end
end

local function match(start, direction)
	local match = {}
	local function recurse(b)
		if not b or b.color ~= start.color then
			return
		end
		table.insert(match, b)
		recurse(b[direction])
	end
	recurse(start)
	mark_match(match)
	return match
end


function M.create()
	local instance = {}
	
	local board = {}
	
	local largest_match = {}
	
	function instance.update(block_ids)
		board = {}
		-- build board
		for _,id in pairs(block_ids) do
			local url = msg.url(nil, id, "script")
			board[hash_to_hex(id)] = { 
				id = id,
				url = url,
				color = go.get(url, "color"),
				left = go.get(url, "left"),
				right = go.get(url, "right"),
				up = go.get(url, "up"),
				down = go.get(url, "down")
			}
		end
		
		-- connect blocks
		for k,v in pairs(board) do
			v.left = board[hash_to_hex(v.left)]
			v.right = board[hash_to_hex(v.right)]
			v.up = board[hash_to_hex(v.up)]
			v.down = board[hash_to_hex(v.down)]
		end
		
		-- find matches
		for _,v in pairs(board) do
			match(v, "up")
			--match(v, "down")
			--match(v, "left")
			match(v, "right")
		end
	end
	
	function instance.largest_match()
		return largest_match
	end
	
	function instance.delete_matches()
		for k,v in pairs(board) do
			if v.match then
				go.delete(v.id)
				board[k] = nil
			end
		end
	end
	
	function instance.has_matches()
		for k,v in pairs(board) do
			if v.match then
				return true
			end
		end
		return false
	end
	
	return instance
end

return M