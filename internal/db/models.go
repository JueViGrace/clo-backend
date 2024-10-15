// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"time"
)

type KeDataconex struct {
	KedCodigo string
	KedNombre string
	KedStatus int32
	KedEnlace string
	KedAgen   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
<<<<<<< HEAD
=======

type KeDoccti struct {
	// Agencia
	Agencia string
	// Tipo de documento
	Tipodoc string
	// numero de documento
	Documento string
	// Tipo de documento visual
	Tipodocv string
	// Código de cliente
	Codcliente string
	// Nombre de Cliente
	Nombrecli string
	// Contribuyente especial
	Contribesp string
	// Cliente en ruta parme
	RutaParme string
	// Tipo de precio aplicado
	Tipoprecio string
	// Emision del documento
	Emision time.Time
	// recepcion de la mercancía
	Recepcion time.Time
	// vencimiento del documento
	Vence time.Time
	// días de crédito
	Diascred string
	// estatus del documento
	Estatusdoc string
	// monto neto del doc
	Dtotneto string
	// Total impuestos
	Dtotimpuest string
	// Total final del documento
	Dtotalfinal string
	// Total monto pagado
	Dtotpagos string
	// Total monto descuentos
	Dtotdescuen string
	// Monto del flete incluido en doc
	Dflete string
	// Monto total de devoluciones
	Dtotdev string
	// Monto total dev x vend
	Dvndmtototal string
	// Monto total de retenciones
	Dretencion string
	// Monto retencion IVA
	Dretencioniva string
	// Código del vendedor
	Vendedor string
	// Código del coordinador
	Codcoord string
	// Ultima actualización del registro
	Fechamodifi time.Time
	// si acepta dev el doc (0 = no 1= si)
	Aceptadev string
	// Si posee negociación especial
	KtiNegesp string
	// IVA del documento en bolívares
	Bsiva string
	// Flete del documento en bolívares
	Bsflete string
	// Retención total del documento en bolívares ?
	Bsretencion string
	// Retención del IVA del documento en bolívares ?
	Bsretencioniva string
	// Tasa del Documento cuando es facturado ?
	Tasadoc string
	// Monto de dscto aplicado en divisas
	Mtodcto string
	// Fecha hasta la cual estará en vigencia el descuento
	Fchvencedcto time.Time
	// Edo del dcto (0=vigente,1=Por aplicar, 2=aplicado, 3=Vencido o Anulado)
	Tienedcto string
	// Lo que hay que cobrar en total de las retenciones del documento en bolívares
	Cbsret string
	// Lo que hay que cobrar en total de las retenciones del documento en dólares
	Cdret string
	// Lo que hay que cobrar de la retención del IVA en bolívares
	Cbsretiva string
	// Lo que hay que cobrar de la retención del IVA en dólares
	Cdretiva string
	// Lo que hay que cobrar de la retención del PARME en bolívares
	Cbsrparme string
	// Lo que hay que cobrar de la retención del PARME en dólares
	Cdrparme string
	// Lo que hay que cobrar de la de la retención del Flete en bolívares
	Cbsretflete string
	// Lo que hay que cobrar de la retención del Flete en dólares
	Cdretflete string
	// Cobrado del IVA en bolívares
	Bsmtoiva string
	// Cobrado del Flete en bolívares
	Bsmtofte string
	// Parme cobrado del documento en bolívares bs
	RetmunMto string
	// Documento con flete dolarizado 0 = no, 1 = si
	Dolarflete int32
	// Lo cobrado de la retención de flete en bolívares BS
	Bsretflete string
	// Lo cobrado de la retención de flete en dólares $
	Dretflete string
	// Parme cobrado del documento en dólares $
	DretmunMto string
	// Retención de IVA Obligatoria
	Retivaoblig uint8
	// 0 = Facturado,
	// 1 = En tránsito,
	// 2 = Entregado
	Edoentrega uint8
	// Moto cobrando del IVA en divisas
	Dmtoiva string
	// % Dcto a aplicado
	Prcdctoaplic string
	// Monto del dcto en $
	Montodctodol string
	// Monto del dcto en Bs
	Montodctobs string
}

type KeDoclmv struct {
	// Agencia
	Agencia string
	// Tipo de documento
	Tipodoc string
	// Numero de documento
	Documento string
	// Tipo de documento visual
	Tipodocv string
	// Grupo del articulo
	Grupo string
	// Subgrupo del articulo
	Subgrupo string
	// Origen del articulo
	Origen string
	// Código del artículo
	Codigo string
	// Código art. dependiente
	Codhijo string
	// Id de linea
	Pid string
	// Nombre de articulo
	Nombre string
	// Cantidad de unds
	Cantidad string
	// Cantidad devuelta
	Cntdevuelt string
	// Cantidad dev x vendedor
	Vndcntdevuelt string
	// Monto total dev x vend
	Dvndmtototal string
	// Precio final en divisas
	Dpreciofin string
	// Precio unitario en divisas
	Dpreciounit string
	// Monto Neto en divisas
	Dmontoneto string
	// Monto total en divisas
	Dmontototal string
	// Tasa de IVA del artículo
	Timpueprc string
	// Tasa de IVA del artículo
	Unidevuelt string
	// Emision del documento
	Fechadoc time.Time
	// Código del vendedor
	Vendedor string
	// Código del coordinador
	Codcoord string
	// Ultima actualización del registro
	Fechamodifi time.Time
}
>>>>>>> 14869e
