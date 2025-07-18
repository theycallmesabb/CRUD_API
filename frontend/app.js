const baseURL = "http://localhost:8000";

function renderForm(type) {
  const formArea = document.getElementById("form-area");
  document.getElementById("task-list").innerHTML = "";

  let formHTML = `<form onsubmit="handle${type}(event)">`;

  if (type === "add" || type === "update") {
    formHTML += `
      <input type="number" id="id" placeholder="Task ID" required />
      <input type="text" id="task" placeholder="Task Description" required />
      <input type="text" id="done" placeholder="Is Done? (true/false)" required />
    `;
  } else if (type === "delete") {
    formHTML += `
      <input type="number" id="id" placeholder="Task ID to Delete" required />
    `;
  }

  formHTML += `<input type="submit" value="${type.charAt(0).toUpperCase() + type.slice(1)} Task" /></form>`;
  formArea.innerHTML = formHTML;
}

async function handleadd(event) {
  event.preventDefault();
  const id = parseInt(document.getElementById("id").value);
  const task = document.getElementById("task").value;
  const done = document.getElementById("done").value.toLowerCase() === "true";

  await fetch(`${baseURL}/add`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ Id: id, Task: task, Done: done })
  });

  alert("Task added!");
  event.target.reset();
}

async function handleupdate(event) {
  event.preventDefault();
  const id = parseInt(document.getElementById("id").value);
  const task = document.getElementById("task").value;
  const done = document.getElementById("done").value.toLowerCase() === "true";

  await fetch(`${baseURL}/update`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ Id: id, Task: task, Done: done })
  });

  alert("Task updated!");
  event.target.reset();
}

async function handledelete(event) {
  event.preventDefault();
  const id = parseInt(document.getElementById("id").value);

  await fetch(`${baseURL}/delete`, {
    method: "DELETE",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ Id: id })
  });

  alert("Task deleted!");
  event.target.reset();
}

async function showTasks() {
  const res = await fetch(`${baseURL}/show`);
  const data = await res.json();
  const list = document.getElementById("task-list");
  const formArea = document.getElementById("form-area");

  formArea.innerHTML = "";
  list.innerHTML = "<h3>📋 Task List:</h3>";
  if (data.length === 0) {
    list.innerHTML += `<p>No tasks available.</p>`;
  } else {
    data.forEach(task => {
      list.innerHTML += `
        <div class="task ${task.Done ? 'done' : ''}">
          <strong>ID:</strong> ${task.Id}<br>
          <strong>Task:</strong> ${task.Task}<br>
          <strong>Done:</strong> ${task.Done}
        </div>
      `;
    });
  }
}
