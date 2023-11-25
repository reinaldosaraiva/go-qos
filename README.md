# Simulação de QoS com Go e Goroutines

## Introdução
Este projeto demonstra a simulação de Qualidade de Serviço (QoS) em redes LAN usando a linguagem de programação Go. Ele consiste em um aplicativo client-server onde o servidor processa mensagens recebidas com diferentes prioridades, demonstrando o conceito de QoS.

## Instalação

Para executar este projeto, você precisa ter Go instalado em seu sistema. [Baixe e instale o Go](https://golang.org/dl/) seguindo as instruções oficiais.

## Uso

Para rodar o servidor, use:

```bash
go run server/main.go
```

Em outro terminal, inicie o cliente:

```bash
go run client/main.go
```

O cliente enviará mensagens aleatórias para o servidor, que irá processá-las, dando prioridade às mensagens marcadas como "PRIORIDADE".

## Licença

Este projeto está licenciado sob a BeerWare
