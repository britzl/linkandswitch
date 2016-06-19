local block = require "linkandswitch.block"

local M = {}

local BLOCK_DISTANCE = block.SIZE + 1
local MAX_LINK_DISTANCE = math.sqrt(BLOCK_DISTANCE * BLOCK_DISTANCE + BLOCK_DISTANCE * BLOCK_DISTANCE)

--- Create a linker instance
-- @param ignore_color Set to tru if the linked should allow links
-- to be created of mixed colors or not (defaults to false)
-- @param max_length Maximum length of links that can be created (defaults
-- to math.huge)
-- @return The linked instance
function M.create(ignore_color, max_length)
	ignore_color = ignore_color or false
	max_length = max_length or math.huge

	local instance = {}
	
	local link = {}
	
	---  add a block to the link
	-- @param id
	function instance.add(id)
		table.insert(link, { id = id, color = block.get_color(id) })
		block.select(id)
	end
	
	--- remove a specific block from the link
	-- @param id
	function instance.remove(id)
		for k,v in pairs(link) do
			if v.id == id then
				link[k] = nil
				block.deselect(id)
				return
			end
		end
	end
	
	--- remove all blocks from the link
	function instance.remove_all()
		for k,v in pairs(link) do
			link[k] = nil
			block.deselect(v.id)
		end
	end
	
	--- check if the link is empty
	-- @return true if link is empty
	function instance.is_empty()
		return #link == 0
	end
	
	--- get the length of the link
	-- @return Length of link
	function instance.length()
		return #link
	end
	
	--- delete all blocks in the link
	function instance.delete()
		for _,v in pairs(link) do
			go.delete(v.id)
		end
		link = {}
	end
	
	--- get the last block in the link, with an optional offset from the last block
	-- @param offset Offset from the last block in the link
	-- @return The block or nil if offset is too large
	function instance.last(offset)
		return link[#link + (offset or 0)]
	end
	
	--- get a specific block in the link
	-- @param index
	-- @return The block or nil if the index was invalid
	function instance.get(index)
		return link[index]
	end
	
	--- try to add a block to the link
	-- A block will only be added to the link if it's adjacent to the last block in
	-- the chain
	-- @param id Id of the block to add
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

	-- check if the link contains a specific block
	-- @param id
	-- @return true if the link contains the block
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