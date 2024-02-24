.PHONY: gencert

prepare:
	mkdir ./cert/${sub}.${domain}

# make gencert day=1 sub=test domain="domain.ru"
gencert: prepare
	openssl req -x509 -nodes -days $(day) -newkey rsa:2048 -keyout ./cert/${sub}.${domain}/priv.key -out ./cert/${sub}.${domain}/cert.crt -subj "/C=RU/ST=RND/L=Rnd/O=${domain}/OU=${domain}/CN=${sub}.${domain}"