local block = require "linkandswitch.block"

local M = {}

local BLOCK_DISTANCE = block.SIZE + 1
local MAX_LINK_DISTANCE = math.sqrt(BLOCK_DISTANCE * BLOCK_DISTANCE + BLOCK_DISTANCE * BLOCK_DISTANCE)

function M.create(ignore_color, max_length)
	ignore_color = ignore_color or false
	max_length = max_length or math.huge

	local instance = {}
	
	local link = {}
	
	function instance.add(id)
		table.insert(link, { id = id, color = block.get_color(id) })
		block.select(id)
	end
	
	function instance.remove(id)
		for k,v in pairs(link) do
			if v.id == id then
				link[k] = nil
				block.deselect(id)
				return
			end
		end
	end
	
	function instance.remove_all()
		for k,v in pairs(link) do
			link[k] = nil
			block.deselect(v.id)
		end
	end
	
	function instance.is_empty()
		return #link == 0
	end
	
	function instance.length()
		return #link
	end
	
	function instance.delete()
		for _,v in pairs(link) do
			go.delete(v.id)
		end
		link = {}
	end
	
	function instance.last(offset)
		return link[#link + (offset or 0)]
	end
	
	function instance.get(index)
		return link[index]
	end
	
	
	function instance.try_to_add(id)
		local last = instance.last()
		local second_to_last = instance.last(-1)
		local distance = last and vmath.length(go.get_position(last.id) - go.get_position(id)) or 0
		if distance <= MAX_LINK_DISTANCE then
			if instance.contains(id) then
				if second_to_last and id == second_to_last.id then
					instance.remove(last.id)
				end
			elseif instance.length() < max_length then
				if not last or block.get_color(id) == last.color or ignore_color then
					instance.add(id)
				end
			end
		end
	end
	
	function instance.dump()
		local s = ""
		for _,l in ipairs(link) do
			s = s .. tostring(l.id) .. " "
		end
		print(s)
	end
	
	function instance.contains(id)
		for _,v in pairs(link) do
			if v.id == id then
				return true
			end
		end
		return false
	end
	
	return instance
end


return M