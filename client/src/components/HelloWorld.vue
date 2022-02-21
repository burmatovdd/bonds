
<template>
  <div class="container">
    <div class="title">Enter Year</div>
    <input id="input" type="text" class="input">
    <button id="button" class="button" @click="sendData">Submit</button>
  </div>
</template>

<script>
"use strict";
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
            const div = document.createElement("div");
            const table = document.createElement("table");
            const thead = document.createElement("thead");
            const thTr = document.createElement("tr");
            const thName = document.createElement("th");
            const thCount = document.createElement("th");
            table.className = "bondsTable"
            div.className = "table-responsive"

            table.appendChild(thead);
            thead.appendChild(thTr);
            thTr.appendChild(thName);
            thTr.appendChild(thCount);

            thName.className = "thName";
            thCount.className = "thCount";

            thName.innerText = "Name";
            thCount.innerText = "Count";

            let a = [];
            const thDate = document.createElement("th");
            thDate.className = "thDate";
            thTr.appendChild(thDate);

            const thDataTrTitle = document.createElement("tr");
            thDate.appendChild(thDataTrTitle);

            const thDataTdTitle = document.createElement("td");
            thDataTrTitle.appendChild(thDataTdTitle);
            thDataTdTitle.innerText = "Date";
            thDataTdTitle.setAttribute("colspan","12")

            const yearTr = document.createElement("tr");
            thDate.appendChild(yearTr);
            for (let i = 1; i <= 12; i++){
              const yearTd = document.createElement("td");
              yearTr.appendChild(yearTd);
              yearTd.className = "yearTd";
              if (i < 10){
                a.push(tmpl.year+"-0"+i);
                yearTd.innerText = tmpl.year+"-0"+i;
              }else{
                a.push(tmpl.year+"-"+i);
                yearTd.innerText = tmpl.year+"-"+i;
              }
            }
            console.log(a);

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

              const bodyTdValue = document.createElement('td');
              bodyTr.appendChild(bodyTdValue);
              bodyTdValue.className = "bodyTdValue";

              const bodyValuesTr = document.createElement('tr');
              bodyTdValue.appendChild(bodyValuesTr)
              for (let j = 0; j < a.length; j++){
                const bodyTrTdValue = document.createElement("td");
                bodyValuesTr.appendChild(bodyTrTdValue);
                let findedCoupon =data.bondInfos[i].coupons.find(coupon=>{
                  console.log("a[j]: ", a[j]);
                  console.log("month: ", coupon.date);
                  return a[j]=== coupon.date
                });
                if (findedCoupon){
                  // console.log("data.bondInfos[i].coupons[j]: ", data.bondInfos[i].coupons[j]);
                  bodyTrTdValue.innerText = findedCoupon.value;
                }
              }
            }
            div.appendChild(table);
            document.body.append(div);
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
