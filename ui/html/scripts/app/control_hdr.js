export default () => ({ 
    testStr: 'Test string',

    onMouseEnter(event) {
        event.preventDefault();
        console.log(`On mouse enter...` + event.target.id);
        let target = event.target;
       
        if (target.id === 'menu1') {
            console.log('menu1 - mouse enter');
            let sib = target.nextElementSibling;
            if (sib) {
                sib.style.display = 'block';
            }
        } else {
            console.log('pop - mouse enter');
            target.style.display = 'block';
        }
    },
    onMouseLeave(event) {
        event.preventDefault();
        console.log(`On mouse leave...` + event.target.id);
        let target = event.target;
        
        if (target.id === 'menu1') {
            console.log('menu1 - mouse leave');
            let sib = target.nextElementSibling;
            if (sib) {
                sib.style.display = 'none';
            }             
        } else {
            console.log('pop - mouse leave');
            target.style.display = 'none';
        }
    },
    onClick(event) {
        console.log(`On click...` + event.target);
        let target = event.target;
        //target.closest(".dropdown_menu").style.display = 'none';
    }
})