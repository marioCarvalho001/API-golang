# API-golang
API de veículos com o objetivo de passar a localização, informações do veiculo e da frota que pertence o veículo. 


 <h1 align="center"> Objetivo e etapas do projeto </h1>
 
 Projeto consiste em uma API de veiculos com 5 endpoints que desenvolvida em Golang e Mysql.
 Para rodar executar o projeto:
  ⏺️  - Tenha o docker instalador na maquina:
          - no terminal digite  ( docker-compose up )
  
  <h3 align="center"> Executar o projeto </h3>
  
  ⏺️  - Tenha o goleng e mysql instalador na maquina: 
          - execute o código da pasta db no terminal do mysql.
          - digite no terminal ( go run main.go )
          
   <h3 align="center"> Testar o Projeto </h3>     
   
  ⏺️ - No projeto tem uma pasta com o nome (Teste)
      - Execute os teste com o uso do Postman seguindo a ordem de get e put com os endpoints e body que estão prontos no arquivo.
  
   <h3 align="center"> Status do projeto </h3>
  
 ✅ Desenvolvido
 
 ❎ Em desenvolvimento
 
 EndPoints, Funcionalidade e informações dos dados.
 
 ⏺️ /fleets  GET/PORT
 ✅- Cadastra a frota 
 ✅- Mysql ( id | name | max_speed )
   - GET: Cadastra a frota 
   ✅ - Deve validar o nome e velocidade do veiculo
   ✅ - a velocidade deve ser maior que 0
   - POST: 
   ✅ - retorna o valor da frota

 
 ⏺️ /fleets/{id}/alerts  
   - Responsavel por enviar um alerta para a frota do veiculo cadastrado quando ultrapassar o limite máximo.
   - Mysql ( id | fleet_id | webhook )
   - GET: 
   ✅ - Retorna o webhook da frota
   - Post um novo webhook  
   ✅ - retorna Code 400 caso a url for invalida
   
   
 ⏺️ /vehicles
   - Contem informações do veiculo e velocidade.
   - Mysql ( id | fleet_id | name | max_speed )
   - GET: retorna o max_speed do veiculo.
    ❎ - Caso não tenha deve ser ultilizar o da frota cadastrado
   - POST: Cadastra um veiculo e sua informações
    ✅  - validar os campos e retorna Code 400 caso um campo seja invalido.
    ✅ - Max_speed deve ser opcional.
 
 ⏺️ /vehicles/{id}/positions
   - Contem informações do veiculo e velocidade.
   - Mysql ( id | vehicle_id | times_stamp | latitude | longitude | current_speed | max_speed )
   - GET: 
     ✅ - lista todas as posições do veiculo cadastrado
   
   - POST: Salva a posição do veiculo 
     ❎ - validar os campos e retorna Code 400 caso um campo seja invalido. 
 
 ⏺️ Emissão de Evento
     ❎ - Após salvar a posição do veículo, você deverá verificar se a velocidade do veículo é maior do que a
           cadastrada (devemos verificar a velocidade cadastrada no veículo e se não tiver, usar a da frota).
           Caso a velocidade seja maior, deverá enviar as seguintes informações para todos os webhooks
           cadastrados da frota:
     ❎ - Se o serviço não responder com um HTTP 200, você deve tentar reenviar 3 vezes, sendo 1, 5 e 15
           segundos após cada tentativa. Se todas as alternativas falharem, você pode cancelar a notificação.

   
