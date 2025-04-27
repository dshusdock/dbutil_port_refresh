import { doc } from "prettier";

export default () => ({ 
    testStr: 'Test string',
    enableClickDetect: true,

    onInit() {
        console.log("My tablevw ...on init");
        document.body.addEventListener("updateMessage", function(event) {
            console.log("Update message event received: ", event.detail);
        });

        document.body.addEventListener("htmx:afterSwap", function(event) {
            console.log("After swap event received: ", event.detail.target.id);

            this.enableClickDetect = true;
        });
    },

    onOutsideHdrClick(event) {
        
        //event.preventDefault();
        
        console.log(`On onOutsideHdrClick click...`);

    },
    onPageCtrlClick(event) {
        event.preventDefault();
        console.log(`On page ctrl click...`);
        let target = event.target;
        var classList = target.classList; // Access the class list
        console.log(target); // Logs a DOMTokenList
        console.log(classList.contains('page_select')); // Check if the element has a specific class
        console.log(`Target: ${target.className}`);
        if (classList.contains('page_select')) {
            console.log('page_select - click');
            let el = document.getElementById('pager_thingy');
            if (el) {
                el.style.display = 'flex';
            }
        }
    },
    onOutsideClick(event) {
        console.log(`On outside click.....`);
        if (this.enableClickDetect) {
            console.log('Running htmx ajax call');
            htmx.ajax('POST', '/event/element/click', {swap: 'innerHTML', target: '#pager_thingy', values: {'eat':'this'}}  ).then(function (response) {
                console.log(response);
                //this.enableClickDetect = false;
            });
           
        }
        
    }

})