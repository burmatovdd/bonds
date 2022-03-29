
<template>
  <div class="container">
    <div class="title">Enter Year</div>
    <input id="input" type="text" class="input">
    <button id="button" class="button" @click="sendData">Submit</button>
  </div>
<!--  <div class="table-responsive">-->
<!--    <table class="bondsTable">-->
<!--      <thead>-->
<!--      <tr>-->
<!--        <th class="thName">Name</th>-->
<!--        <th class="thCount">Count</th>-->
<!--        <th class="thDate">-->
<!--          <table>-->
<!--            <tr>-->
<!--              <td colspan="12">Date</td>-->
<!--            </tr>-->
<!--          </table>-->
<!--          <table class="months">-->
<!--            <tr id="months" v-for="month in monthArray" :key="month.day">-->
<!--              {{month.day}}-->
<!--            </tr>-->
<!--          </table>-->
<!--        </th>-->
<!--      </tr>-->
<!--      </thead>-->
<!--    </table>-->
<!--  </div>-->
</template>

<script>
import { defineComponent } from 'vue';
export default defineComponent({
  setup(){
    let tmpl = {
      year : null,
    }
    // let monthArray = []
    async function sendData (){
      tmpl.year = document.getElementById("input").value

      let sendUrl = "http://localhost:8080/year";

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
            // monthArray = [
            //   {day: tmpl.year + "-" + "01"},
            //   {day: tmpl.year + "-" + "02"},
            //   {day: tmpl.year + "-" + "03"},
            //   {day: tmpl.year + "-" + "04"},
            //   {day: tmpl.year + "-" + "05"},
            //   {day: tmpl.year + "-" + "06"},
            //   {day: tmpl.year + "-" + "07"},
            //   {day: tmpl.year + "-" + "08"},
            //   {day: tmpl.year + "-" + "09"},
            //   {day: tmpl.year + "-" + "10"},
            //   {day: tmpl.year + "-" + "11"},
            //   {day: tmpl.year + "-" + "12"},
            // ]
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

            const ThDataTrTable = document.createElement("table");
            const thDataTrTitle = document.createElement("tr");

            ThDataTrTable.appendChild(thDataTrTitle)
            thDate.appendChild(ThDataTrTable);

            const thDataTdTitle = document.createElement("td");
            thDataTrTitle.appendChild(thDataTdTitle);
            thDataTdTitle.innerText = "Date";
            thDataTdTitle.setAttribute("colspan","12")

            const yearTable = document.createElement("table");
            const yearTr = document.createElement("tr");

            yearTable.appendChild(yearTr)
            thDate.appendChild(yearTable);

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

            const tBody = document.createElement("tbody");
            table.appendChild(tBody);

            let yearSum = 0;
            for (let i = 0; i < data.allInfos.bondInfos.length;i++){
              const bodyTr = document.createElement("tr");
              bodyTr.className = "bodyTr";
              tBody.appendChild(bodyTr);

              const bodyTdName = document.createElement("td");
              bodyTr.appendChild(bodyTdName);
              bodyTdName.innerText = data.allInfos.bondInfos[i].bond.name;

              const bodyTdCount = document.createElement('td');
              bodyTr.appendChild(bodyTdCount);
              bodyTdCount.className = "bodyTdCount";
              bodyTdCount.innerText = data.allInfos.bondInfos[i].bond.count;

              const bodyTdValue = document.createElement('td');
              bodyTr.appendChild(bodyTdValue);
              bodyTdValue.className = "bodyTdValue";

              const bodyValuesTable = document.createElement('table');
              const bodyValuesTr = document.createElement('tr');

              bodyValuesTable.appendChild(bodyValuesTr)
              bodyTdValue.appendChild(bodyValuesTable)

              for (let j = 0; j < a.length; j++){
                const bodyTrTdValue = document.createElement("td");

                bodyValuesTr.appendChild(bodyTrTdValue);
                bodyTrTdValue.className = "bodyTrTdValue";
                //  алгоритм заполнения данных по месяцам
                let findedCoupon = data.allInfos.bondInfos[i].coupons.find(coupon=>{
                  return a[j] === coupon.date
                });
                if (findedCoupon){
                  bodyTrTdValue.innerText = findedCoupon.value;
                  yearSum += findedCoupon.value;
                }
              }
            }

            const newYearSum = yearSum.toFixed();// округление числа
            console.log("yearSum: ", newYearSum.toString());
            const divYearSum = document.createElement("div");
            divYearSum.className = "divYearSum"
            divYearSum.innerText = "Year sum: " + newYearSum.toString() + " " + "rub";

            const divMonthValueContainer = document.createElement("div");
            divMonthValueContainer.className = "divMonthValueContainer";
            for ( let i = 0; i < data.allInfos.months.length; i++){

              const newMonthValue = data.allInfos.months[i].value.toFixed();
              const divMonthValue = document.createElement("div");
              divMonthValue.innerText = "Month " + data.allInfos.months[i].date + " :" + newMonthValue.toString() + " " + "rub";
              divMonthValueContainer.appendChild(divMonthValue);

            }

            div.appendChild(table);

            document.body.append(div);
            document.body.append(divYearSum);
            document.body.append(divMonthValueContainer);

          }).catch(console.error);
    }
    return{
      tmpl,
      sendData,
      // monthArray
    }
  }
})
</script>

<style lang="css">
@import "../../style/style.css";
@import "./inputYear.css";
</style>
