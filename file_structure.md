foolio@foolio-HP-Laptop-14-dq0xxx:/media/foolio/beans/emv_system$ ls -R
.:
client  database  engine  schemas  server

./client:
package.json  src  vite.config.ts

./client/src:
generated  network  state  ui  visualization

./client/src/generated:
emv_protocol_generated.ts

./client/src/network:
binary_codec.ts  websocket_client.ts

./client/src/state:
autocomplete_store.ts  card_tree_store.ts

./client/src/ui:
AutocompleteSearch.tsx  TagInspector.tsx  TerminalView.tsx

./client/src/visualization:
canvas_stage.ts  layout_calculator.ts  node_renderer.ts

./database:
db.go  migrations  queries

./database/migrations:
0001_create_emv_tables.sql  0002_add_trigram_indexes.sql

./database/queries:
files.sql  records.sql  tags.sql

./engine:
apdu  emv  reader  tlv

./engine/apdu:
command.go  response.go  status_words.go

./engine/emv:
dictionary.go  hierarchy.go  validator.go

./engine/reader:
pcsc_driver.go  session.go

./engine/tlv:
decoder.go  length.go  tag.go  value.go

./schemas:
emv_protocol.fbs

./server:
handlers  main.go  transport

./server/handlers:
apdu_handler.go  query_handler.go  stream_handler.go

./server/transport:
client_connection.go  router.go  websocket_hub.go
foolio@foolio-HP-Laptop-14-dq0xxx:/media/foolio/beans/emv_system$ 



