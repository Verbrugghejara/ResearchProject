<template>
  <!-- <div class="w-11 overflow-auto h-11"
    v-on:scroll.passive="scrollOutput($event)"
  >
    <p>this is scrollable</p>
    <p>this is scrollable</p>
  </div> -->
  <div
    class="h-screen overflow-auto"
    v-on:scroll.passive="scrollOutput($event)"
    ref="scrollable"
  >
    <div class="flex justify-center">
      <p
        class="absolute bottom-10 rounded-2xl bg-red-400 p-2 px-6 text-5xl shadow-lg"
        v-if="scrollData > 0"
      >
        {{ scrollData}}
      </p>
    </div>
    <div class="h-[1080px] border-b-2 bg-red-500 text-9xl font-bold">
      <h1 class="pt-5 text-center text-6xl font-semibold">App 1</h1>
      <p class="py-[18%] text-center">1</p>
    </div>
    <div class="h-[1080px] border-b-2 bg-red-500 text-9xl font-bold">
      <p class="py-[25%] text-center">2</p>
    </div>
    <div class="h-[1080px] border-b-2 bg-red-500 text-9xl font-bold">
      <p class="py-[25%] text-center">3</p>
    </div>

    <!-- <p v-for="i in Array.from({ length: 50 }, (v, i) => i)">
      {{ i }}
    </p> -->
  </div>
</template>

<script lang="ts">
import { watch } from "fs";
import { computed, defineComponent, onMounted, ref, Ref } from "vue";
export default defineComponent({
  name: "App",
  setup() {
    // let scrollData = -1;
    let mainEvent: any;
    
    const scrollData: Ref<number> = ref(-1)

    async function funcName(this: any, url: string) {
      fetch(url)
        .then((response) => response.json())
        .then((data) => {
          console.log(data)
          scrollData.value = data.Data;
          mainEvent.target.scrollTop = scrollData.value; 
        });
    }
    funcName("http://localhost:3333/api/63d28d3cfa172fc285b7a85f");
    window.addEventListener("click", (event) => {
      mainEvent = event;
      // @ts-ignore
      event.target.scrollTop = scrollData;
    });

    const scrollOutput = (e: any | undefined) => {
      console.log(e.target.scrollTop);
      scrollData.value = e.target.scrollTop;
      fetch("http://localhost:3333/api/update/63d28d3cfa172fc285b7a85f", {
        method: "PUT",
        headers: {
          "Content-type": "application/json",
        },
        body: JSON.stringify({ Data: e.target.scrollTop }),
      });
      return e.target.scrollTop;
    };

    return {
      scrollOutput,
      scrollData,
    };
  },
  mounted() {
    console.log("mounted");
    // @ts-ignore
    this.$refs.scrollable.click();
  },
  watch: {
    scrollData: function (val) {
      // console.log("Scroll data changed");
    },
  },
});
</script>
