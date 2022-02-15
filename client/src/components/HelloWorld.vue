
<template>
  <div class="container">
    <div class="title">Enter Year</div>
    <input id="input" type="text" class="input">
    <button id="button" class="button" @click="sendData">Submit</button>
  </div>
</template>

<script>
import { defineComponent } from 'vue';
export default defineComponent({
  setup(){
    let tmpl = {
      year : null,
    }
    async function sendData (){
      tmpl.year = document.getElementById("input").value

      let sendUrl = "http://localhost:8080/result";

      await fetch(sendUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(tmpl)
      }).then((response) => {
        return response.json();
      })
          .then((data) => {
            console.log(data);

            const table = document.createElement("table");
            const thead = document.createElement("thead");
            const thTr = document.createElement("tr");
            const thName = document.createElement("th");
            const thCount = document.createElement("th");
            const thDate = document.createElement("th");
            const thValue = document.createElement("th");
            table.className = "bondsTable"

            table.appendChild(thead);
            thead.appendChild(thTr);
            thTr.appendChild(thName);
            thTr.appendChild(thCount);
            thTr.appendChild(thDate);
            thTr.appendChild(thValue);

            thName.className = "thName";
            thCount.className = "thCount";
            thDate.className = "thDate";
            thValue.className = "thValue";

            thName.innerText = "Name";
            thCount.innerText = "Count";
            thDate.innerText = "Date";
            thValue.innerText = "Value";

            const tBody = document.createElement("tbody");
            table.appendChild(tBody);

            for (let i = 0; i < data.bondInfos.length;i++){
              const bodyTr = document.createElement("tr");
              tBody.appendChild(bodyTr);

              // bodyTr.className = "name tableText"
              const bodyTdName = document.createElement("td");
              bodyTr.appendChild(bodyTdName);
              bodyTdName.innerText = data.bondInfos[i].bond.name;

              const bodyTdCount = document.createElement('td');
              bodyTr.appendChild(bodyTdCount);
              bodyTdCount.className = "bodyTdCount";
              bodyTdCount.innerText = data.bondInfos[i].bond.count;

              const bodyTdDate = document.createElement("td");
              bodyTr.appendChild(bodyTdDate);

              const bodyTdValue = document.createElement('td');
              bodyTr.appendChild(bodyTdValue);
              bodyTdValue.className = "bodyTdValue";

              for(let j = 0; j < data.bondInfos[i].coupons.length;j++){

                const bodyTrDate = document.createElement("tr");
                bodyTdDate.appendChild(bodyTrDate);

                const bodyTrTdDate = document.createElement("td");
                bodyTrDate.appendChild(bodyTrTdDate);
                bodyTrTdDate.innerText = data.bondInfos[i].coupons[j].date;

                const bodyTrValue = document.createElement("tr");
                bodyTdValue.appendChild(bodyTrValue);

                const bodyTrTdValue = document.createElement("td");
                bodyTrValue.appendChild(bodyTrTdValue);
                bodyTrTdValue.innerText = data.bondInfos[i].coupons[j].value;
              }
            }
            document.body.append(table);
          }).catch(console.error);
    }
    return{
      tmpl,
      sendData,
    }
  }
})
</script>

<style lang="css">
@import "../style/style.css";
</style>
