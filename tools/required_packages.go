package tools

func Packages() string {
	return `
		shared_pkgs="libffi-dev python-dev"
		declare -A osInfo;

		osInfo[/etc/redhat-release]=yum
		osInfo[/etc/arch-release]=pacman
		osInfo[/etc/gentoo-release]=emerge
		osInfo[/etc/debian_version]=apt-get
		osInfo[/etc/alpine-release]=apk

		for f in ${!osInfo[@]}
		do
			if [[ -f $f ]];then
				if [ "${osInfo[$f]}" == "yum" ]; then
					sudo yum update && sudo yum install -y $shared_pkgs opus opus-tools ffmpeg
				elif [ "${osInfo[$f]}" == "pacman" ]; then
					sudo pacman -Syu && sudo pacman -S $shared_pkgs opus opus-tools base-devel youtube-dl
				elif [ "${osInfo[$f]}" == "emerge" ]; then
					sudo emerge --update --deep --with-bdeps=y @world && sudo emerge -pv $shared_pkgs opus opus-tools ffmpeg
				elif [ "${osInfo[$f]}" == "apt-get" ]; then
					sudo apt-get upgrade -y && sudo apt-get install -y libopus-dev opus-tools libcurl4-openssl-dev $shared_pkgs ffmpeg build-essential autoconf automake libtool m4 youtube-dl
				elif [ "${osInfo[$f]}" == "apk" ]; then
					sudo apk update && sudo apk add libsodium ffmpeg $shared_pkgs opus opus-tools autoconf automake libtool m4 youtube-dl curl-dev
				fi
			fi
		done
	`
}
