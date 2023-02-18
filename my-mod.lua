local textOffset = { x = 1, y = 1 }
local textPosition = { x = 200, y = 60 }
local text = "VyacheArt"

function calculateDimensions(text)
    local bounds = BoundString(DefaultFont, text)
    local textDimensions = {}
    textDimensions.width = bounds:Dx()
    textDimensions.height = bounds:Dy()

    return textDimensions
end

local textDimensions = calculateDimensions(text)


function draw(screen)
    local backgroundColor = ColorRGBA()
    backgroundColor.R = 10
    backgroundColor.G = 70
    backgroundColor.B = 90
    backgroundColor.A = 255
    screen:Fill(backgroundColor)

    DrawText(screen, text, DefaultFont, textPosition.x, textPosition.y, ColorWhite)
end

function update()
    applyTextOffset()
end

function applyTextOffset()
    textPosition.x = textPosition.x + textOffset.x
    textPosition.y = textPosition.y + textOffset.y

    local screenSize = GetWindowSize()
    if textPosition.x + textDimensions.width > screenSize.X or textPosition.x < 0 then
        textOffset.x = -textOffset.x
    end

    if textPosition.y + textDimensions.height > screenSize.Y or textPosition.y < 0 then
        textOffset.y = -textOffset.y
    end
end
