<template>
  <div class="bg-white shadow overflow-hidden sm:rounded-md">
    <div class="bg-white px-4 py-5 border-b border-gray-200 sm:px-6">
      <div
        class="-ml-4 -mt-2 flex items-center justify-between flex-wrap sm:flex-no-wrap"
      >
        <div class="ml-4 mt-2">
          <h3 class="text-lg leading-6 font-medium text-gray-900">
            All Tracks
          </h3>
        </div>
        <div class="ml-4 mt-2 flex-shrink-0">
          <span class="inline-flex rounded-md shadow-sm">
            <button
              type="button"
              class="relative inline-flex items-center px-4 py-2 border border-transparent text-sm leading-5 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:shadow-outline-indigo focus:border-indigo-700 active:bg-indigo-700"
            >
              Create new track
            </button>
          </span>
        </div>
      </div>
    </div>
    <ul>
      <li v-for="track in tracks" :key="track.id">
        <a
          href="#"
          class="block hover:bg-gray-50 focus:outline-none focus:bg-gray-50 transition duration-150 ease-in-out"
        >
          <div class="px-4 py-4 sm:px-6">
            <div class="flex items-center justify-between">
              <div
                class="text-sm leading-5 font-medium text-indigo-600 truncate"
              >
                <nuxt-link :to="`/tracks/` + track.id">{{
                  track.name
                }}</nuxt-link>
              </div>
              <div class="ml-2 flex-shrink-0 flex">
                <span
                  class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800"
                >
                  {{ track.durationMs | millisecondsToMinutesAndSeconds }}
                </span>
              </div>
            </div>
            <div class="sm:flex">
              <div
                class="mr-6 flex items-center text-sm leading-5 text-gray-500"
              >
                <svg
                  class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                  fill="currentColor"
                  viewBox="0 0 20 20"
                >
                  <path
                    d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3zM6 8a2 2 0 11-4 0 2 2 0 014 0zM16 18v-3a5.972 5.972 0 00-.75-2.906A3.005 3.005 0 0119 15v3h-3zM4.75 12.094A5.973 5.973 0 004 15v3H1v-3a3 3 0 013.75-2.906z"
                  ></path>
                </svg>

                {{ track.artists | valuesAsList("name") }}
              </div>
              <div
                v-if="track.albums"
                class="mr-6 flex items-center text-sm leading-5 text-gray-500"
              >
                <svg
                  class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                  fill="currentColor"
                  viewBox="0 0 20 20"
                >
                  <path
                    d="M7 3a1 1 0 000 2h6a1 1 0 100-2H7zM4 7a1 1 0 011-1h10a1 1 0 110 2H5a1 1 0 01-1-1zM2 11a2 2 0 012-2h12a2 2 0 012 2v4a2 2 0 01-2 2H4a2 2 0 01-2-2v-4z"
                  ></path>
                </svg>

                {{ track.albums | valuesAsList("name") }}
              </div>
              <div
                class="mr-6 flex items-center text-sm leading-5 text-gray-500"
              >
                <svg
                  class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                  fill="currentColor"
                  viewBox="0 0 20 20"
                >
                  <path
                    fill-rule="evenodd"
                    d="M17.707 9.293a1 1 0 010 1.414l-7 7a1 1 0 01-1.414 0l-7-7A.997.997 0 012 10V5a3 3 0 013-3h5c.256 0 .512.098.707.293l7 7zM5 6a1 1 0 100-2 1 1 0 000 2z"
                    clip-rule="evenodd"
                  ></path>
                </svg>

                {{ track.genres | valuesAsList("name") }}
              </div>
            </div>
          </div>
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      tracks: []
    };
  },

  head() {
    return {
      title: "Tracks"
    };
  },

  async asyncData({ app, error }) {
    const query = `
      query tracks {
        tracks {
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
      }`;

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
      tracks: data.data.tracks
    };
  }
};
</script>

<style></style>
