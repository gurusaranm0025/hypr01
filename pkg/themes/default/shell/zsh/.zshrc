# Path to your oh-my-zsh installation.
export ZSH="$HOME/.oh-my-zsh"
source $ZSH/oh-my-zsh.sh

# See https://github.com/ohmyzsh/ohmyzsh/wiki/Themes
ZSH_THEME="lukerandall"
# ZSH_THEME="common"

plugins=(
  git
  zsh-syntax-highlighting
  zsh-autosuggestions
)

# rust bins
export PATH="$HOME/.cargo/bin:$PATH"

# auto suggest
ZSH_AUTOSUGGEST_HIGHLIGHT_STYLE="fg=#663399,standout"

# nvm config
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && . "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && . "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

zstyle ':completion:*' special-dirs true
setopt autocd

export PATH="$HOME/.local/bin/:$HOME/.config/npm/bin:$PATH" # Tailscale Systray binary is in .lcoal/bin

export XDG_STATE_HOME="$HOME/.local/state"
export XDG_SRC_HOME="$HOME/.local/src"
export XDG_DATA_HOME="$HOME/.local/share"
export XDG_CONFIG_HOME="$HOME/.config"
export XDG_CACHE_HOME="$HOME/.cache"

export LESS="--ignore-case"

export CARGO_HOME="$HOME/.config/rust-cargo"
export NPM_CONFIG_USERCONFIG="$XDG_CONFIG_HOME/npm/npmrc"
export GTK2_RC_FILES="$XDG_CONFIG_HOME/gtk-2.0/gtkrc":"$XDG_CONFIG_HOME/gtk-2.0/gtkrc.mine"

export GNUPGHOME="$XDG_DATA_HOME/gnupg"
export PYTHON_HISTORY="$XDG_STATE_HOME/python_history"
export _JAVA_OPTIONS=-Djava.util.prefs.userRoot="$XDG_CONFIG_HOME"/java

export CUDA_CACHE_PATH="$XDG_CACHE_HOME/cuda"
export __GL_SHADER_DISK_CACHE=1
export __GL_SHADER_DISK_CACHE_PATH="$XDG_CACHE_HOME/nv"
export __GL_SHADER_DISK_CACHE_SKIP_CLEANUP=1
export __GL_THREADED_OPTIMIZATIONS=0

# eza -al format ouput to diference permissions
export EZA_COLORS="ur=32:uw=32:ux=32:gr=36:gw=36:gx=36:tr=31:tw=31:tx=31"

# ALIASES
alias ls='eza -a --git --icons=automatic --group-directories-first'
alias ll='eza -al'
alias lt='eza -a --tree --level=2 --group-directories-first --icons=automatic'

alias cat="bat"
alias ccat="/bin/cat"

alias vim-zshrc="vim ~/.zshrc"
alias vimzshrc="vim $XDG_CONFIG_HOME/zsh/.zshrc"
alias vimsudo='sudo -E nvim /etc/sudoers'

alias wget='wget --hsts-file="$XDG_CACHE_HOME/wget-hsts"'
alias g='grep'
alias man="COLUMNS=100 man"

alias mkdir="mkdir -p"
alias mv='mv -v'

alias search-package="pacman -Q | grep -i" # search through installed packages.
alias search-pacman="pacman -Ss" # search through pacman repo
alias search-font="fc-list | grep -i" # search through available font styles.
alias search-env="env | grep -i" # search through environment variables.
alias search-port="sudo ss -tulpn | grep -i"

alias refresh-fonts="fc-cache -fv" # refreshed font cache
alias refresh-shell="source ~/.zshrc" # source .zshrc after changes

alias system-update="sudo pacman -Syu --noconfirm && yay -Syu && flatpak update"

# Set-up FZF key bindings (CTRL R for fuzzy history finder)
source <(fzf --zsh)

HISTFILE=~/.config/zsh/.zsh_history
HISTSIZE=1000
SAVEHIST=1000
setopt appendhistory

# NEW EDITS

# export KEEPER_STORAGE_DIR="$XDG_STATE_HOME/keeper_storage"
# export WGETRC="$XDG_CONFIG_HOME/wgetrc"
# export MYSQL_HISTFILE="$XDG_DATA_HOME/mysql/history"
# export DOCKER_CONFIG="$XDG_CONFIG_HOME/docker"

