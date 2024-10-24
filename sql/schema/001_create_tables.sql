-- +goose Up
CREATE TABLE IF NOT EXISTS ke_dataconex (
  id CHAR(36) NOT NULL PRIMARY KEY,
  ked_codigo VARCHAR(6) NOT NULL UNIQUE,
  ked_nombre VARCHAR(80) NOT NULL DEFAULT '',
  ked_status BOOLEAN NOT NULL DEFAULT 0,
  ked_enlace TEXT NOT NULL,
  ked_agen VARCHAR(20) NOT NULL DEFAULT '',
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS ke_wcnf_conf (
  id CHAR(36) NOT NULL PRIMARY KEY,
  cnfg_idconfig char(30) NOT NULL,
  cnfg_clase char(1) NOT NULL DEFAULT '',
  cnfg_tipo char(1) NOT NULL,
  cnfg_valnum DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  cnfg_valsino BOOLEAN NOT NULL DEFAULT 0,
  cnfg_valtxt TEXT NOT NULL,
  cnfg_lentxt SMALLINT NOT NULL DEFAULT 0,
  cnfg_valfch date NOT NULL DEFAULT '1000-01-01',
  cnfg_activa BOOLEAN NOT NULL DEFAULT 0,
  cnfg_etiq TEXT NOT NULL,
  cnfg_ttip TEXT NOT NULL,
  user_id CHAR(36) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NULL,
  CONSTRAINT UC_wcnf_idconfig UNIQUE (cnfg_idconfig, user_id)
);

CREATE TABLE IF NOT EXISTS articulo (
  id CHAR(36) NOT NULL PRIMARY KEY,
  codigo VARCHAR(25) NOT NULL UNIQUE,
  grupo INT NOT NULL DEFAULT 0,
  subgrupo INT NOT NULL DEFAULT 0,
  nombre VARCHAR(150) NOT NULL DEFAULT '',
  referencia VARCHAR(20) NOT NULL DEFAULT '',
  marca VARCHAR(20) NOT NULL DEFAULT '',
  unidad VARCHAR(15) NOT NULL DEFAULT '',
  existencia INT NOT NULL DEFAULT 0,
  precio1 DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  precio2 DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  precio3 DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  precio4 DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  precio5 DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  precio6 DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  precio7 DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  discont BOOLEAN NOT NULL DEFAULT 0,
  vta_max INT NOT NULL DEFAULT 0,
  vta_min INT NOT NULL DEFAULT 0,
  dctotope DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  enpreventa BOOLEAN NOT NULL DEFAULT 0,
  comprometido INT NOT NULL DEFAULT 0,
  vta_minenx INT NOT NULL DEFAULT 0,
  vta_solofac BOOLEAN NOT NULL DEFAULT 0,
  vta_solone BOOLEAN NOT NULL DEFAULT 0,
  codbarras VARCHAR(20) NOT NULL DEFAULT '',
  detalles TEXT NOT NULL,
  cantbulto INT NOT NULL DEFAULT 0,
  costo_prom DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  util1 DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  util2 DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  util3 DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  fchultcomp DATE NOT NULL DEFAULT '1000-01-01',
  qtyultcomp INT NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS usuario (
  id CHAR(36) NOT NULL PRIMARY KEY,
  username varchar(30) NOT NULL UNIQUE DEFAULT '',
  password varchar(20) NOT NULL DEFAULT '',
  role ENUM('cliente', 'vendedor', 'gerente', 'administrador') NOT NULL DEFAULT 'cliente',
  desactivo BOOLEAN NOT NULL DEFAULT 0,
  ult_sinc TIMESTAMP NOT NULL DEFAULT NOW(),
  version varchar(30)  NOT NULL DEFAULT '0.0.0',
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE ke_nivgcia (
  id CHAR(36) NOT NULL PRIMARY KEY,
  kng_codgcia varchar(8) NOT NULL DEFAULT '',
  kng_codcoord varchar(8) NOT NULL DEFAULT '',
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS vendedor ( 
  id CHAR(36) NOT NULL PRIMARY KEY,
  user_id CHAR(36) DEFAULT NULL,
  codigo VARCHAR(8) NOT NULL UNIQUE DEFAULT '',
  nombre VARCHAR(60) NOT NULL DEFAULT '',
  telefono_1 VARCHAR(30) NOT NULL DEFAULT '',
  telefono_2 VARCHAR(30) NOT NULL DEFAULT '',
  telefono_movil VARCHAR(100) NOT NULL DEFAULT '',
  status SMALLINT NOT NULL DEFAULT 0,
  supervpor VARCHAR(8) NOT NULL DEFAULT 0,
  sector INT NOT NULL DEFAULT 0,
  subcodigo INT NOT NULL DEFAULT 0,
  email TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS cliente (
  id CHAR(36) NOT NULL PRIMARY KEY,
  user_id CHAR(36) DEFAULT NULL,
  codigo VARCHAR(20) NOT NULL UNIQUE,
  nombre VARCHAR(100) NOT NULL DEFAULT '',
  direccion VARCHAR(200) NOT NULL DEFAULT '',
  telefonos VARCHAR(150) NOT NULL DEFAULT '',
  perscont VARCHAR(50) NOT NULL DEFAULT '',
  vendedor VARCHAR(8) NOT NULL DEFAULT '',
  contribespecial BOOLEAN NOT NULL DEFAULT 0,
  status SMALLINT NOT NULL DEFAULT 0,
  sector INT NOT NULL DEFAULT 0,
  subcodigo INT NOT NULL DEFAULT 0,
  precio SMALLINT NOT NULL DEFAULT 1,
  email TEXT NOT NULL,
  kne_activa BOOLEAN NOT NULL DEFAULT 0,
  kne_mtomin DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  noemifac BOOLEAN NOT NULL DEFAULT 0,
  noeminota BOOLEAN NOT NULL DEFAULT 0,
  fchultvta DATE NOT NULL DEFAULT '1000-01-01',
  mtoultvta DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  prcdpagdia DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  promdiasp DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  riesgocrd DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  cantdocs INT NOT NULL DEFAULT 0,
  totmtodocs DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  prommtodoc DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  diasultvta DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  promdiasvta DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  limcred DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  dolarflete BOOLEAN NOT NULL DEFAULT 0,
  nodolarflete BOOLEAN NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS ke_estadc01 (
  id CHAR(36) NOT NULL PRIMARY KEY,
  codcoord VARCHAR(8) NOT NULL DEFAULT '',
  nomcoord VARCHAR(54) NOT NULL DEFAULT '',
  vendedor VARCHAR(8) NOT NULL UNIQUE,
  nombrevend VARCHAR(54) NOT NULL DEFAULT '',
  cntpedidos INT NOT NULL DEFAULT 0,
  mtopedidos DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  cntfacturas INT NOT NULL DEFAULT 0,
  mtofacturas DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  metavend DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  prcmeta DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  cntclientes INT NOT NULL DEFAULT 0,
  clivisit INT NOT NULL DEFAULT 0,
  prcvisitas DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  lom_montovtas DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  lom_prcvtas DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  lom_prcvisit DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  rlom_montovtas DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  rlom_prcvtas DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  rlom_prcvisit DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  fecha_estad date NOT NULL DEFAULT '1000-01-01',
  ppgdol_totneto DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  devdol_totneto DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  defdol_totneto DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  totdolcob DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  cntrecl DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  mtorecl DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS ke_opti (
  id CHAR(36) NOT NULL PRIMARY KEY,
  kti_ndoc VARCHAR(17) NOT NULL UNIQUE,
  kti_tdoc VARCHAR(3) NOT NULL DEFAULT '',
  kti_codcli VARCHAR(20) NOT NULL DEFAULT '',
  kti_nombrecli VARCHAR(100) NOT NULL DEFAULT ' ',
  kti_codven VARCHAR(8) NOT NULL DEFAULT '',
  kti_docsol CHAR(1) NOT NULL DEFAULT '',
  kti_condicion CHAR(1) NOT NULL DEFAULT '',
  kti_tipprec CHAR(1) NOT NULL DEFAULT '0',
  kti_totneto DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  kti_status CHAR(1) NOT NULL DEFAULT '0',
  kti_nroped VARCHAR(8) NOT NULL DEFAULT '',
  kti_fchdoc DATETIME NOT NULL DEFAULT '1000-01-01 00:00:00',
  kti_negesp BOOLEAN NOT NULL DEFAULT 0,
  ke_pedstatus CHAR(2) NOT NULL DEFAULT '',
  dolarflete BOOLEAN NOT NULL DEFAULT 0,
  complemento BOOLEAN NOT NULL DEFAULT 0,
  nro_complemento VARCHAR(17) NOT NULL DEFAULT '',
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS ke_opmv (
  pedido_id CHAR(36) NOT NULL,
  articulo_id CHAR(36) NOT NULL,
  kti_tdoc VARCHAR(8) NOT NULL DEFAULT '',
  kti_ndoc VARCHAR(17) NOT NULL,
  kti_tipprec CHAR(1) NOT NULL DEFAULT '1',
  kmv_codart VARCHAR(25) NOT NULL,
  kmv_nombre VARCHAR(150) NOT NULL DEFAULT '',
  kmv_cant INT NOT NULL DEFAULT 0.00,
  kmv_artprec DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  kmv_stot DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  kmv_dctolin DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NULL,
  CONSTRAINT PK_pedido_lineas PRIMARY KEY (pedido_id, articulo_id),
  CONSTRAINT UC_ndoc_lineas UNIQUE (kti_ndoc, kmv_codart)
);

CREATE TABLE IF NOT EXISTS ke_doccti (
  id CHAR(36) NOT NULL PRIMARY KEY,
  agencia VARCHAR(3) NOT NULL DEFAULT '',
  tipodoc VARCHAR(3) NOT NULL DEFAULT '',
  documento VARCHAR(8) NOT NULL UNIQUE,
  tipodocv VARCHAR(3) NOT NULL DEFAULT '',
  codcliente VARCHAR(20) NOT NULL DEFAULT '',
  nombrecli VARCHAR(100) NOT NULL DEFAULT '',
  contribesp BOOLEAN NOT NULL DEFAULT 0,
  ruta_parme BOOLEAN NOT NULL DEFAULT 0,
  tipoprecio CHAR(1) NOT NULL DEFAULT '1',
  emision DATE NOT NULL DEFAULT '1000-01-01',
  recepcion DATE NOT NULL DEFAULT '1000-01-01',
  vence DATE NOT NULL DEFAULT '1000-01-01',
  diascred DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  estatusdoc CHAR(1) NOT NULL DEFAULT '0',
  dtotneto DECIMAL(8,4) NOT NULL DEFAULT 0.00,
  dtotimpuest DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  dtotalfinal DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  dtotpagos DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  dtotdescuen DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  dFlete DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  dtotdev DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  dvndmtototal DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  dretencion DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  dretencioniva DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  vendedor VARCHAR(8) NOT NULL DEFAULT '',
  codcoord VARCHAR(8) NOT NULL DEFAULT '',
  aceptadev BOOLEAN NOT NULL DEFAULT 1,
  kti_negesp BOOLEAN NOT NULL DEFAULT 0,
  bsiva DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  bsflete DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  bsretencion DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  bsretencioniva DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  tasadoc DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  mtodcto DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  fchvencedcto DATE NOT NULL DEFAULT '1000-01-01',
  tienedcto BOOLEAN NOT NULL DEFAULT 0,
  cbsret DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  cdret DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  cbsretiva DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  cdretiva DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  cbsrparme DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  cdrparme DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  cbsretflete DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  cdretflete DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  bsmtoiva DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  bsmtofte DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  retmun_mto DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  dolarflete BOOLEAN NOT NULL DEFAULT 0,
  bsretflete DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  dretflete DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  dretmun_mto DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  retivaoblig BOOLEAN NOT NULL DEFAULT 0,
  edoentrega BOOLEAN NOT NULL DEFAULT 0,
  dmtoiva DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  prcdctoaplic DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  montodctodol DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  montodctobs DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS ke_doclmv (
  doc_id CHAR(36) NOT NULL,
  articulo_id CHAR(36) NOT NULL,
  agencia VARCHAR(3) NOT NULL DEFAULT '',
  tipodoc VARCHAR(3) NOT NULL DEFAULT '',
  documento VARCHAR(8) NOT NULL,
  tipodocv VARCHAR(3) NOT NULL DEFAULT '',
  grupo INT NOT NULL DEFAULT 0,
  subgrupo INT NOT NULL DEFAULT 0,
  origen DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  codigo VARCHAR(25) NOT NULL,
  codhijo VARCHAR(25) NOT NULL DEFAULT '',
  pid VARCHAR(12) NOT NULL DEFAULT '',
  nombre VARCHAR(100) NOT NULL DEFAULT '',
  cantidad INT NOT NULL DEFAULT 0,
  cntdevuelt INT NOT NULL DEFAULT 0,
  vndcntdevuelt DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  dvndmtototal DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  dpreciofin DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  dpreciounit DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  dmontoneto DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  dmontototal DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  timpueprc DECIMAL(8, 4) NOT NULL DEFAULT 0.00,
  unidevuelt INT NOT NULL DEFAULT 0,
  fechadoc DATE NOT NULL DEFAULT '1000-01-01',
  vendedor VARCHAR(8) NOT NULL DEFAULT '',
  codcoord VARCHAR(8) NOT NULL DEFAULT '',
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NULL,
  CONSTRAINT PK_docc_lineas PRIMARY KEY (doc_id, articulo_id),
  CONSTRAINT UC_docc_lineas UNIQUE (documento, codigo)
);

-- +goose Down
DROP TABLE ke_dataconex;
DROP TABLE ke_wcnf_conf;
DROP TABLE articulo;
DROP TABLE usuario;
DROP TABLE vendedor;
DROP TABLE cliente;
DROP TABLE ke_estadc01;
DROP TABLE ke_opti;
DROP TABLE ke_opmv;
DROP TABLE ke_doccti;
DROP TABLE ke_doclmv;
DROP TABLE ke_nivgcia;

