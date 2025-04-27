export default () => ({ 
    testStr: '',
    prevElement: null,
    showUserDropdown: false,
    lightMode: false,
    onInit() {
        console.log("header ...on init");
      
    },
    openUserDropdown(event) {
        console.log("Open dropdown");
        event.preventDefault();
        //document.getElementsByClassName("header-dropdown")[0].classList.toggle("show");
        this.showUserDropdown = !this.showUserDropdown;
    },
    toggleColorMode(event) {
        console.log("Toggle color mode");
        const root = document.documentElement;
        if (this.lightMode) {
            this.lightMode = false;
            root.style.setProperty('--background-pri-color', 'var(--dark-color)');
            root.style.setProperty('--background-sec-color', 'var(--dark-color)');
            root.style.setProperty('--background-ter-color', 'var(--dark-color)');
        }
        else {
            this.lightMode = true;
            root.style.setProperty('--background-pri-color', 'var(--light-color)');
        }
    },
    openAppSidePanel(event) {
        document.getElementsByClassName("appsidenavvw")[0].classList.toggle("closeappsidenav");
    }
     
})