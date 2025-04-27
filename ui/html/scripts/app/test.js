export default () => ({ 
    testStr: 'Test string',

    onTestClick(event) {
        event.preventDefault();
        
        console.log(`On test click...`);
        console.log(`Test string: ${this.testStr}`);
    },
})