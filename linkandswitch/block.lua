local M = {}

M.SIZE = 106

function M.get_color(id)
	return go.get(msg.url(nil, id, "script"), "color")
end

function M.select(id)
	msg.post(id, "select")
end

function M.deselect(id)
	msg.post(id, "deselect")
end

return M