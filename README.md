# API-golang
API de veículos com o objetivo de passar a localização, informações do veiculo e da frota que pertence o veículo. 


 <h1 align="center"> Objetivo e etapas do projeto </h1>
 
 Projeto consiste em uma API de veiculos com 5 endpoints que foram desenvolvida em Golang e Mysql.
 
 <h3 align="center"> Executar o projeto </h3>
 
 
 Para rodar executar o projeto:<br>
  ⏺️  - Se tiver o docker instalador na maquina:<br>
          - no terminal digite  ( docker-compose up )<br>
  <br>
  ⏺️  - Se tiver o goleng e mysql instalador na maquina: <br>
          - execute o código da pasta db no terminal do mysql.<br>
          - digite no terminal ( go run main.go )<br>
   <br>       
   <h3 align="center"> Testar o Projeto </h3>     
   
  ⏺️ - No projeto tem uma pasta com o nome (Teste)<br>
      - Execute os teste com o uso do Postman seguindo a ordem de put e get com os endpoints e body que estão prontos no arquivo.
  
   <h3 align="center"> Status do projeto </h3>
  
 ✅ Desenvolvido <br>
 
 ❎ Em desenvolvimento <br>
 
 EndPoints, Funcionalidade e informações dos dados. <br>
 
 ⏺️ /fleets  GET/PORT <br>
 ✅- Cadastra a frota  <br>
 ✅- Mysql ( id | name | max_speed ) <br>
   - GET: Cadastra a frota  <br>
   ✅ - Deve validar o nome e velocidade do veiculo <br>
   ✅ - a velocidade deve ser maior que 0 <br>
   - POST: 
   ✅ - retorna o valor da frota <br>

 
 ⏺️ /fleets/{id}/alerts  <br>
   - Responsavel por enviar um alerta para a frota do veiculo cadastrado quando ultrapassar o limite máximo. <br>
   - Mysql ( id | fleet_id | webhook ) <br>
   - GET:  <br>
   ✅ - Retorna o webhook da frota <br>
   - Post um novo webhook   <br>
   ✅ - retorna Code 400 caso a url for invalida <br>
   
   
 ⏺️ /vehicles <br>
   - Contem informações do veiculo e velocidade. <br>
   - Mysql ( id | fleet_id | name | max_speed ) <br>
   - GET: retorna o max_speed do veiculo. <br>
    ❎ - Caso não tenha deve ser ultilizar o da frota cadastrado <br>
   - POST: Cadastra um veiculo e sua informações <br>
    ✅  - validar os campos e retorna Code 400 caso um campo seja invalido. <br>
    ✅ - Max_speed deve ser opcional. <br>
 
 ⏺️ /vehicles/{id}/positions <br>
   - Contem informações do veiculo e velocidade. <br>
   - Mysql ( id | vehicle_id | times_stamp | latitude | longitude | current_speed | max_speed ) <br>
   - GET:   <br>
     ✅ - lista todas as posições do veiculo cadastrado <br>
   
   - POST: Salva a posição do veiculo  <br>
     ❎ - validar os campos e retorna Code 400 caso um campo seja invalido.  <br>
 
 ⏺️ Emissão de Evento <br>
     ❎ - Após salvar a posição do veículo, você deverá verificar se a velocidade do veículo é maior do que a
           cadastrada (devemos verificar a velocidade cadastrada no veículo e se não tiver, usar a da frota).
           Caso a velocidade seja maior, deverá enviar as seguintes informações para todos os webhooks
           cadastrados da frota:  <br>
     ❎ - Se o serviço não responder com um HTTP 200, você deve tentar reenviar 3 vezes, sendo 1, 5 e 15
           segundos após cada tentativa. Se todas as alternativas falharem, você pode cancelar a notificação.  <br>

   
