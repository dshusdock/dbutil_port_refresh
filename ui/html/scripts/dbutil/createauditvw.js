export default () => ({ 
    testStr: 'Test string',

    onTestClick(event) {
        event.preventDefault();
        
        console.log(`On test click...`);
        console.log(`Test string: ${this.testStr}`);
    },
    onCloseClick(event) {
        console.log("on close clicked");
        event.preventDefault();
  
        document.getElementById('blur-overlay').style.display = 'none';
        document.getElementById('create_auditvw').style.display = 'none';
    }
})