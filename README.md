# Live Stream Server

Esse projeto é um projeto de estudos onde desenvolvi o sistema de live-stream utilizando as seguintes tecnologias:

- Golang
- Nginx
- Postgres
- Docker

## Auth Server

O Auth Server é um servidor de autenticação escrito em Go que recebe uma chave de autenticação de um programa de stream (OBS, Streamlabs) referente a uma live-stream que está sendo executada e verifica se essa live-stream está no banco de dados.

Se der match com algum registro do banco, o programa de stream starta a live, caso contrário, um erro é exibido no programa e nos logs da aplicação.

## NGINX

Usamos o nginx para levantar um servidor rtmp para transmissões ao vivo. A imagem usada foi a `tiangolo/nginx-rtmp` que pode ser acessada aqui: [nginx-rtmp](https://github.com/tiangolo/nginx-rtmp-docker)

## Playback Server

O Playback Server é um servidor que irá consumir a live stream gerada pelo servidor nginx, pormitindo o usuário acessar determinada url pelo navegador e consumir a live.

## PostgresSQL

Este banco de dados tem o único propósito de guardar as chaves de live-stream.


## Executando o projeto.

Para executar o projeto, você precisará do docker e de um programa de live stream na sua máquina.

- Execute o comando `docker compose up --build` para rodar todos os servidores acima.
- Depois que todos os containers estiverem de pé, abra seu programa de live-stream, acesse as configurações e coloque a chave de transmissão (que pode ser encontrada no arquivo `scripts/2_stream_keys_populate.sql`). 
A chave de transmissão é composta por: `nomedalive_streamkey`. Ex: `primeiralive_569a7708-1f77-4ccb-9be3-e1acf8706a68` É importante que haja essa separação por underscore.
- Inicie a sua transmissão.
- No navegador, acesse: `http://localhost:8001/live/:live/index.m3u8` (em :live você precisa colocar o nome da live que está sendo iniciada, ex: `primeiralive`)

Desta forma o seu projeto já estará funcionando e você poderá consumir a sua própria transmissão pelo navegador.
