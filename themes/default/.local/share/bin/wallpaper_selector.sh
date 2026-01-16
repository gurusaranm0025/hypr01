#!/bin/bash

IMAGES_DIR="$HOME/.HyprOne/Walls"
CACHE_DIR="$HOME/.cache/HyprOne/Walls"
CURRENT_WALL_DIR="$HOME/.config/hyprone/current_wall"

mkdir -p "$CACHE_DIR"
shopt -s nullglob

for img in "$IMAGES_DIR"/*.{png,jpg,jpeg,gif,webp}; do
    [ -f "$img" ] || continue
    thumb="$CACHE_DIR/$(basename "$img").png"
    if [ ! -f "$thumb" ]; then
        magick "${img[0]}" -thumbnail 220x220^ -gravity center -extent 500x400^ "$thumb"
    fi
    entry="$(basename "$img")\x00icon\x1f$img"
    options+=("$entry")
done 

selection="$(echo -e "$(printf '%s\n' "${options[@]}")" | rofi -dmenu -show-icons -i -theme "$HOME/.config/rofi/themes/wallpapers.rasi")"

echo "$selection"

if [ -n "$selection" ] && [ -f "$IMAGES_DIR/$selection" ]; then
	cp "$IMAGES_DIR/$selection" "$CURRENT_WALL_DIR/wallpaper"
	swww img "$CURRENT_WALL_DIR/wallpaper" --transition-type fade --transition-duration 0.5
fi
