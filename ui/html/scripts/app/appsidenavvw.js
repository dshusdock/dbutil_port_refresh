export default () => ({ 
    testStr: 'Test string',
    

    onCloseSideNavClick() {
        document.getElementsByClassName("appsidenavvw")[0].classList.toggle("closeappsidenav");
    },
})