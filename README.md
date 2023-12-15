# estacionamiento_car

Programación concurrente

Diseño e implementación de una aplicación informática multihilos.

Utilizar el mecanismo de semáforos para los casos de exclusión mutua y
condiciones de sincronización.

Estacionamiento

Implementar una aplicación concurrente utilizando como mecanismo de
sincronización semáforos en Go.

Requerimientos obligatorios:
· - Deberá simularse la creación de 100 vehículos(Llegadas a través de distribución poison) o de manera infinita.
· - El estacionamiento deberá tener capacidad para 20 vehículos.


Conforme los vehículos requieran el servicio, deberán verificar si el estacionamiento tiene
espacios disponibles en ese momento, generándose dos posibles decisiones:
a) Si no hay lugar disponible, el vehículo en cuestión se deberá bloquear; estará en ese estado,
hasta que un vehículo abandone el estacionamiento y libere a los vehículos bloqueados.
b) Si hay lugares disponibles, deberá intentar acceder al lugar. La entrada principal al
estacionamiento sirve tanto de entrada como de salida, es decir, es un recurso compartido
por los vehículos que ingresen y que salen del lugar. Generándose las siguientes decisiones:
a. Si la puerta de acceso está ocupada por un vehículo que abandona el estacionamiento,
el vehículo que intenta entrar deberá bloquearse, y se desbloqueará hasta que el
recurso compartido este disponible.
b. Si la puerta de acceso esta ocupada por un vehículo que está entrando, este vehículo
también podrá entrar, y todos los que tengan el mismo sentido.
c. En los casos de vehículos que están saliendo, aplica la regla anterior.

Finalmente, cuando haya ingresado al lugar deberá verificar si los cajones de
estacionamiento están disponibles para que, pueda ocuparlos; cuando verifique a
estos, se podrán tomar las siguientes decisiones:
a) Si el cajón en turno está disponible, el vehículo lo ocupará un tiempo aleatorio
entre 1 y 5 seg, tras lo cual lo deberá desocupar.
b) Si el cajón está ocupado, deberá irse al siguiente cajón, y así sucesivamente
hasta encontrar un cajón disponible.
