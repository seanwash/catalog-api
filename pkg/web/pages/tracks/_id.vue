<template>
  <div class="bg-white shadow overflow-hidden  sm:rounded-lg">
    <div class="bg-white px-4 py-5 border-b border-gray-200 sm:px-6">
      <div
        class="-ml-4 -mt-2 flex items-center justify-between flex-wrap sm:flex-no-wrap"
      >
        <div class="ml-4 mt-2">
          <h3 class="text-lg leading-6 font-medium text-gray-900">
            Track Details
          </h3>
        </div>
        <div class="ml-4 mt-2 flex-shrink-0">
          <span class="inline-flex rounded-md shadow-sm">
            <nuxt-link
              to="/"
              class="relative inline-flex items-center px-4 py-2 border border-transparent text-sm leading-5 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:shadow-outline-indigo focus:border-indigo-700 active:bg-indigo-700"
            >
              Back to Tracks
            </nuxt-link>
          </span>
        </div>
      </div>
    </div>
    <div class="px-4 py-5 sm:p-0">
      <dl>
        <div class="sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
          <dt class="text-sm leading-5 font-medium text-gray-500">
            Name
          </dt>
          <dd
            class="mt-1 text-sm leading-5 text-gray-900 sm:mt-0 sm:col-span-2"
          >
            {{ track.name }}
          </dd>
        </div>
        <div
          class="mt-8 sm:mt-0 sm:grid sm:grid-cols-3 sm:gap-4 sm:border-t sm:border-gray-200 sm:px-6 sm:py-5"
        >
          <dt class="text-sm leading-5 font-medium text-gray-500">
            Written by
          </dt>
          <dd
            class="mt-1 text-sm leading-5 text-gray-900 sm:mt-0 sm:col-span-2"
          >
            {{ track.artists | valuesAsList("name") }}
          </dd>
        </div>
        <div
          class="mt-8 sm:mt-0 sm:grid sm:grid-cols-3 sm:gap-4 sm:border-t sm:border-gray-200 sm:px-6 sm:py-5"
        >
          <dt class="text-sm leading-5 font-medium text-gray-500">
            Appears on
          </dt>
          <dd
            class="mt-1 text-sm leading-5 text-gray-900 sm:mt-0 sm:col-span-2"
          >
            {{ track.albums | valuesAsList("name") }}
          </dd>
        </div>
        <div
          class="mt-8 sm:mt-0 sm:grid sm:grid-cols-3 sm:gap-4 sm:border-t sm:border-gray-200 sm:px-6 sm:py-5"
        >
          <dt class="text-sm leading-5 font-medium text-gray-500">
            Genres
          </dt>
          <dd
            class="mt-1 text-sm leading-5 text-gray-900 sm:mt-0 sm:col-span-2"
          >
            {{ track.genres | valuesAsList("name") }}
          </dd>
        </div>
        <div
          class="mt-8 sm:mt-0 sm:grid sm:grid-cols-3 sm:gap-4 sm:border-t sm:border-gray-200 sm:px-6 sm:py-5"
        >
          <dt class="text-sm leading-5 font-medium text-gray-500">
            Duration
          </dt>
          <dd
            class="mt-1 text-sm leading-5 text-gray-900 sm:mt-0 sm:col-span-2"
          >
            {{ track.durationMs | millisecondsToMinutesAndSeconds }}
          </dd>
        </div>
      </dl>
    </div>
  </div></template
>

<script>
export default {
  data() {
    return {
      track: null
    };
  },

  head() {
    return {
      title: this.track.name
    };
  },

  async asyncData({ app, params, error }) {
    const id = params.id;

    const query = `
      query track {
        track(id: ${id}) {
          id
          name
          durationMs
          artists {
            id
            name
          }
          albums {
            id
            name
          }
          genres {
            id
            name
          }
        }
      }
    `;

    const { data, errors } = await app.$axios.post(
      `${process.env.API_URL}/graphql`,
      { query: query }
    );

    if (errors && errors.length) {
      return error({
        statusCode: 500,
        message: "An error occurred fetching tracks."
      });
    }

    return {
      track: data.data.track
    };
  }
};
</script>

<style scoped></style>
