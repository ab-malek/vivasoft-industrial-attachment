const todolist = [];

renderHtml();


function renderHtml(){
    let todoList = '';
    console.log(todolist);
    for(let i = 0;i < todolist.length; i++){
        todoList += `
        <div class="list-grid">
            <div class="child">${todolist[i].name}</div>
            <div class="child"> ${todolist[i].date} </div>
            <button onclick="
                todolist.splice(${i},1);
                renderHtml();
            " class="delete-button child">Delete</button>
        </div>
        `;
    }
    document.querySelector('.js-todo-list').innerHTML = todoList;
}


function addTodo(){
    const inputElement = document.querySelector('.js-input-name');
    const dueName = inputElement.value;
    if(dueName === ''){
        alert('Please write your todo');
        return;
    }
    const inputDateElement = document.querySelector('.js-input-date');
    const dueDate = inputDateElement.value;
    todolist.push({
        name : dueName,
        date : dueDate
    });
    renderHtml();
    inputElement.value = '';
    inputDateElement.value = '';
}


