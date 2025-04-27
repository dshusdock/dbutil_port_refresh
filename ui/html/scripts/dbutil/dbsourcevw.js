import { doc } from "prettier";

export default () => ({ 
    current: null,
    open: false,
    itemHighlight: null,

    init() {
        console.log("dbsource ...on init");
        document.body.addEventListener("htmx:afterSwap", function(event) {
            console.log("After swap event received: ", event.detail.target.id);
            let text = document.getElementById('db-avail-state').textContent;

            if (text.includes("Available")) {
                console.log("Available");
                document.getElementById('db-source-icon').style.color = "green";
            } else {
                console.log("Not Available");
                document.getElementById('db-source-icon').style.color = "red";
            }
        });
    },
    onClose() {
       let el = document.getElementsByClassName('dbsourcevw')[0];
       if (el) {
           el.style.display = 'none';
       }
    },
    
})