# Vim-Man Level Planning

## Level 1: Basic Movement

*   **Vim Concepts:** `h` (left), `j` (down), `k` (up), `l` (right).
*   **Gameplay:** A simple maze. Player learns to navigate Vim-Man using the basic movement keys to reach an exit tile. Introduce the concept of "Vim" as the source of power/commands.
*   **Visuals:** Clean, straightforward. Maybe a terminal-like aesthetic.
*   **Objective:** Reach the exit.

## Level 2: Exiting Vim (The First Challenge)

*   **Vim Concepts:** `:` (command mode), `q` (quit), `!` (force - for `q!`).
*   **Gameplay:** Player reaches what seems like an exit, but it's blocked. A hint appears: "Type :q to exit". If the player has "unsaved changes" (e.g., collected an item they weren't supposed to, or stepped on a 'dirty' tile), `:q` fails. They then learn `:q!` to force quit (exit the level).
*   **Visuals:** The exit could be a flashing cursor or a representation of a closing window. "Dirty" tiles could have a distinct visual.
*   **Objective:** Successfully exit the level using the correct quit command.

## Level 3: Basic Text Editing (Deletion)

*   **Vim Concepts:** `x` (delete character under cursor), `dw` (delete word - maybe simplified to just deleting a 'word' entity).
*   **Gameplay:** The path is blocked by "error characters" (like `x` tiles) or "bug words" (small enemy-like entities). Player must use `x` to delete single error characters or `dw` (or a simplified version) to remove "bug words" to clear the path.
*   **Visuals:** "Error characters" could be red `x`'s. "Bug words" could be small, distinct sprites.
*   **Objective:** Clear the path and reach the exit.

## Level 4: The Bomberman (Introduction to Insert Mode & Esc)

*   **Vim Concepts:** `i` (insert mode), `Esc` (return to normal mode), placing "bombs" (characters) in insert mode.
*   **Gameplay:** Inspired by Bomberman. Player can enter insert mode with `i`. While in insert mode, pressing a movement key (or a specific "bomb" key) places a "character bomb" that explodes after a short delay, clearing destructible blocks. Player must use `Esc` to return to normal mode to move safely.
*   **Visuals:** Destructible blocks, bomb sprites, explosion effects.
*   **Objective:** Clear a path through destructible blocks and reach the exit.

## Level 5: The Word Jumper's Gauntlet

*   **Vim Concepts:** `w` (next word), `b` (previous word), `e` (end of word), `ge` (end of previous word).
*   **Gameplay:** The level is a series of platforms (words) separated by gaps. Player must use `w`, `b`, `e`, `ge` to jump precisely between platforms. Some platforms might be crumbling, requiring quick successive jumps. Enemies could patrol on longer "sentence" platforms.
*   **Visuals:** Platforms clearly look like "words". Gaps are obvious. Crumbling platforms could animate.
*   **Objective:** Navigate the platforms and reach the exit.

## Level 6: The Repetition Realm (Numeric Precision)

*   **Vim Concepts:** Using numbers with motion/commands (e.g., `3w`, `5j`, `2dd`).
*   **Gameplay:** Paths are blocked by gates requiring a specific number of actions. For example, a gate "3" high requires `3j` to pass if it's a vertical jump, or defeating an enemy might take `2x` (two delete actions). Collectibles could be arranged in patterns that are easiest to get with numbered movements.
*   **Visuals:** Gates could display the number required. Collectibles in clear numerical patterns.
*   **Objective:** Use numbered commands to overcome obstacles and reach the exit.

## Level 7: The Search & Find Expedition

*   **Vim Concepts:** `/` (search forward), `?` (search backward), `n` (next occurrence), `N` (previous occurrence).
*   **Gameplay:** Key items or "rescue targets" (e.g., specific characters) are hidden across a maze-like map. Using `/` or `?` followed by the target character would highlight the path or the item. Enemies might change characters, forcing the player to re-search.
*   **Visuals:** Maze environment. Searched items/paths could glow or have a special indicator.
*   **Objective:** Find the target(s) using search commands and reach the exit.

## Level 8: The Copy-Paste Assembly Line

*   **Vim Concepts:** `y` (yank/copy), `p` (paste after), `P` (paste before), Visual mode (`v` character-wise, `V` line-wise) for selecting text.
*   **Gameplay:** The player needs to construct a specific "word" or "code snippet" by yanking parts from different areas of the level and pasting them into a designated "build zone." Visual mode could be used to select multi-character "components."
*   **Visuals:** Distinct "source" areas for yanking. A clear "build zone" where pasting occurs. Visual feedback for selection in Visual mode.
*   **Objective:** Assemble the target word/snippet and complete the level. 