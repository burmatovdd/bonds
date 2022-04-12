
<template>
  <div class="container">
    <div class="title">Enter Year</div>
    <input id="input" type="text" class="input">
    <button id="button" class="button" @click="sendData">Submit</button>
  </div>
  <div class="table-responsive">
    <table class="bondsTable">
      <thead>
      <tr>
        <th class="thName">Name</th>
        <th class="thCount">Count</th>
        <th class="thDate">
          <table>
            <tr>
              <td colspan="12">Date</td>
            </tr>
          </table>
          <table class="months">
            <tr id="months" v-for="month in monthArray" :key="month.day" class="months-day">
              {{month.day}}
            </tr>
          </table>
        </th>
      </tr>
      </thead>
      <tbody>
      <tr id="bonds" v-for="bond in bondsArray"  class="bondsTr">
        <td>
          {{bond.name}}
        </td>
        <td>
          {{bond.count}}
        </td>
        <td>
          <table>
            <tr>
              <td v-for="value in valueArray" v-if="findedCoupon">
                {{value.value}}
              </td>
            </tr>
          </table>
        </td>
        <td>
          <button class="button" @click="deleteBond(bond)">Delete</button>
        </td>
      </tr>
      <tr class="bondsTr">
        <td>Total</td>
        <td></td>
        <td>
          <table class="table-total">
            <tr>
              <td v-for="value in totalArray" class="total-value">
                {{value.value}}
              </td>
            </tr>
          </table>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import {defineComponent, ref} from 'vue';
export default defineComponent({
  setup(){
    let tmpl = {
      year : null,
    }
    let bondName = {
      name: null
    }
    const monthArray = ref([]);
    const bondsArray = ref([]);
    const valueArray = ref([]);
    const totalArray = ref([]);
    let findedCoupon = ref()

    async function deleteBond(bond) {
      let sendUrl = "http://localhost:8080/delete";
      bondName.name = bond.name
      for (let i = 0; i < bondsArray.value.length; i++){
        if (bondsArray.value[i].name === bondName.name){
          bondsArray.value.splice(i,1)
        }
      }
      await fetch(sendUrl,{
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(bondName)
      })
    }

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
            monthArray.value = [
              {day: tmpl.year + "-" + "01"},
              {day: tmpl.year + "-" + "02"},
              {day: tmpl.year + "-" + "03"},
              {day: tmpl.year + "-" + "04"},
              {day: tmpl.year + "-" + "05"},
              {day: tmpl.year + "-" + "06"},
              {day: tmpl.year + "-" + "07"},
              {day: tmpl.year + "-" + "08"},
              {day: tmpl.year + "-" + "09"},
              {day: tmpl.year + "-" + "10"},
              {day: tmpl.year + "-" + "11"},
              {day: tmpl.year + "-" + "12"},
            ]

            for (let i = 0; i < data.allInfos.bondInfos.length; i++ ){
              bondsArray.value.push({name: data.allInfos.bondInfos[i].bond.name, count: data.allInfos.bondInfos[i].bond.count})
              for (let j = 0; j < monthArray.value.length; j++){
                findedCoupon = data.allInfos.bondInfos[i].coupons.find(coupon=>{
                  return monthArray[j] === coupon.date
                })
              }
              for (let k = 0; k < data.allInfos.bondInfos[i].coupons.length; k++){
                valueArray.value.push({value: data.allInfos.bondInfos[i].coupons[k].value})
              }
            }

            for (let i = 0; i < data.allInfos.months.length; i++){
              totalArray.value.push({value: data.allInfos.months[i].value.toFixed(2)})
            }

            console.log(data);

          }).catch(console.error);
    }
    return{
      tmpl,
      deleteBond,
      sendData,
      monthArray,
      bondsArray,
      valueArray,
      findedCoupon,
      totalArray
    }
  }
})
</script>

<style lang="css">
@import "../../style/style.css";
@import "./inputYear.css";
</style>
