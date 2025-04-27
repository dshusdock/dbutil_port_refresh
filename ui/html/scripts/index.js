import Alpine from 'alpinejs'
import test from './app/test'
import header from './app/header'
import control_hdr from './app/control_hdr'
import tablevw from './app/tablevw'
import appsidenavvw from './app/appsidenavvw'
import dbsidenavvw from './dbutil/dbsidenavvw'
import dbsourcevw from './dbutil/dbsourcevw'
import auditvw from './dbutil/auditvw'
import createauditvw from './dbutil/createauditvw'

Alpine.data('test', test)
Alpine.data('header', header)
Alpine.data('cntrl_hdr', control_hdr)
Alpine.data('tablevw', tablevw)
Alpine.data('appsidenavvw', appsidenavvw)
Alpine.data('dbsidenavvw', dbsidenavvw)
Alpine.data('dbsourcevw', dbsourcevw)
Alpine.data('auditvw', auditvw)
Alpine.data('createauditvw', createauditvw)

window.Alpine = Alpine

Alpine.start()
//htmx.config.disableInheritance = true;

console.log("Release the hounds!!......again!!...");


