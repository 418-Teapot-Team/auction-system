import axios from 'axios';
export default defineStore('main', () => {
  const config = useRuntimeConfig();
  const auth = useAuthState();
  async function createAuction(payload) {
    try {
      const token = auth.token.value;
      const data = new FormData();
      data.append('auction', JSON.stringify(payload));
      const res = await axios.post(
        config.public.baseAPI + '/auction/create',
        data,
        {
          headers: {
            Authorization: token,
            'Content-Type': 'multipart/form-data',
          },
        }
      );
      return res.data;
    } catch (e) {
      throw new Error(e.message);
    }
  }

  return {
    createAuction,
  };
});
