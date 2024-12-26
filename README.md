# Cep-Race

Este projeto implementa uma solução em GoLang que utiliza multithreading para fazer requisições simultâneas a duas APIs diferentes de consulta de CEP e retorna o resultado mais rápido. O objetivo é otimizar a busca por informações de endereço, garantindo que a resposta seja recebida em até 1 segundo.

## Requisitos

### Descrição

O projeto realiza duas requisições HTTP simultâneas para consultar o endereço de um determinado CEP, usando as seguintes APIs:

- **BrasilAPI**: https://brasilapi.com.br/api/cep/v1/{cep}
- **ViaCEP**: http://viacep.com.br/ws/{cep}/json/

A resposta mais rápida será exibida, junto com o nome da API que retornou a informação.

### Funcionalidades

- Realiza requisições simultâneas para as duas APIs.
- Exibe os dados do endereço e a qual API pertence a resposta mais rápida.
- Limita o tempo de resposta a 1 segundo. Caso contrário, um erro de timeout será exibido.

### Requisitos de Implementação

- **Multithreading**: Utiliza goroutines para realizar as requisições simultâneas.
- **Timeout**: Implementa um timeout de 1 segundo para as requisições.
- **Exibição do Resultado**: Exibe os dados do endereço, juntamente com o nome da API que retornou a resposta mais rápida.
- **Fallback em caso de erro de timeout**: Caso as duas APIs não respondam em 1 segundo, exibe uma mensagem de erro.
