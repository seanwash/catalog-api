import Vue from "vue";

Vue.filter("capitalize", val => val.toUpperCase());

Vue.filter("valuesAsList", (values, key) => {
  if (values && values.length) {
    return values.map(value => value[key]).join(", ");
  } else {
    return "";
  }
});

Vue.filter("millisecondsToMinutesAndSeconds", value => {
  const minutes = Math.floor(value / 60000);
  const seconds = ((value % 60000) / 1000).toFixed(0);
  return minutes + ":" + (seconds < 10 ? "0" : "") + seconds;
});
