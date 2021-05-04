# Nubank dev test

#### Dependências

* Docker

#### Execução

* Compilar o projeto de acordo com sua arquitetura, `make build-linux` para linux, ou `make build-darwin` para macOS.
* Executar o binário gerado na pasta bin de acordo com sua arquitetura `./bin/authorize-linux-amd64` para linux, ou `./bin/authorize-darwin-amd64` para macOS.
* Para outros comandos, consulte `make help`.

#### Notas

* Como não possuo macOS, não garanto o funcionamento nesta plataforma, tendo em vista que dependendo do nível de restrição, o binário não poderá ser executado por ser de "fonte desconhecida".
