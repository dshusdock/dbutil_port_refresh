export default () => ({ 
    testStr: 'Test string',

    onTestClick(event) {
        event.preventDefault();
        
        console.log(`On test click...`);
        console.log(`Test string: ${this.testStr}`);
    },
    checkCheckbox(event) {
        console.log("event: ", event);
        let target = event.target;
        console.log("target: ", target);
        console.log("target.id: ", target.id);
        console.log("target.checked: ", target.checked);
    

    
    },
})