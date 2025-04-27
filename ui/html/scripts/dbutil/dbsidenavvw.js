import { doc } from "prettier";

export default () => ({ 
    testStr: 'Test string',
    prevPanel: null,
    prevChevron: null,
    drawerClosed: true,

    onCatClick(event) {
        console.log(`On cat click...`);
        var panel = event.target.parentElement.nextElementSibling;
        var chevron = event.target.parentElement.lastElementChild;

        if (this.prevPanel != null && this.prevPanel != panel) {
            this.prevPanel.classList.remove("show-it");
            this.prevChevron.classList.remove("rotate-fwd");
        }

        chevron.classList.toggle("rotate-fwd");

        setTimeout((el)=>{
            console.log(`Delayed function...`);
            el.classList.toggle("show-it");  
        }, 300, panel);
        this.prevPanel = panel;
        this.prevChevron = chevron;
    },
    onMenuClick(event) {
        console.log(`On menu click...`);
        var menu = document.getElementsByClassName("dbsidenavvw")[0];
        //let icon = event.target.childNodes[1];
        let icon = document.getElementById("dbsidenav-chevron-icon");
        menu.classList.toggle("slide-in");

        if (this.drawerClosed) {
            this.drawerClosed = false;
            icon.classList.remove("fa-chevron-right");
            icon.classList.add("fa-chevron-left");
        } else {
            this.drawerClosed = true;
            icon.classList.add("fa-chevron-right");
            icon.classList.remove("fa-chevron-left");
        }

    },
    onItemClick(event) {
        console.log(`On item click...`);
        var menu = document.getElementsByClassName("dbsidenavvw")[0];
        let icon = document.getElementById("dbsidenav-chevron-icon");
       
        menu.classList.toggle("slide-in");
    }
})


