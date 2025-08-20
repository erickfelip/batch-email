# Batch-email

### Domain:
Consiste em um sistema de centralizador de envio de e-mails em lote através de campanhas.

#### Features:
- Endpoint para envio dos e-mails em lote com os contatos
- Endpoint para informar o status do lote enviado.
- Os e-mails enviados poderão ser personalizados por contato, com isso, informações do contato poderão estar no texto do e-mail.

### Problema: 
Em um cenário repetitivo de multiplas aplicações que devem lidar com envio de e-mail, existe o processo cansativo de sempre criar esse tipo de serviço, o intuito do projeto é evitar esse retrabalho em aplicações existentes ou novas que precisam desse tipo de serviço de envio de email, com um sistema centralizador de envio de e-mails em lote abstraindo o tipo de envio em campanhas esse processo custo de retrabalho é evitado, pois, o usuario apenas irá criar uma campanha > essa campanha será redirecionada para o lote informado (lista de emails) > acompanhamento do status do lote enviado > os emails enviados poderão ser personalizados por contato ou contexto das aplicações.
