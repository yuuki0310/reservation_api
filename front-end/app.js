const app = Vue.createApp({
  data() {
    return {
      storeId: "",
      year: null,
      month: null,
      storeReservations: null,
      userUuid: "",
      userReservations: null,
      newReservation: {
        uuid: "",
        store_id: null,
        from: "",
        to: "",
      },
      reservationMessage: "",
      availableTimeSlots: [
        { from: "00:00", to: "01:45" },
        { from: "02:00", to: "03:45" },
        { from: "04:00", to: "05:45" },
        { from: "06:00", to: "07:45" },
        { from: "08:00", to: "09:45" },
        { from: "10:00", to: "11:45" },
        { from: "12:00", to: "13:45" },
        { from: "14:00", to: "15:45" },
        { from: "16:00", to: "17:45" },
        { from: "18:00", to: "19:45" },
        { from: "20:00", to: "21:45" },
        { from: "22:00", to: "23:45" },
      ],
    };
  },
  methods: {
    async getStoreReservations() {
      try {
        const response = await fetch(
          `https://8ib4rtllwk.execute-api.ap-northeast-1.amazonaws.com/stores/${this.storeId}/reservations?year=${this.year}&month=${this.month}`
        );
        if (!response.ok)
          throw new Error("店舗の予約状況を取得できませんでした");
        this.storeReservations = await response.json();
      } catch (error) {
        alert(error.message);
      }
    },
    async getUserReservations() {
      try {
        const response = await fetch(
          `https://8ib4rtllwk.execute-api.ap-northeast-1.amazonaws.com/users/${this.userUuid}/reservations`
        );
        if (!response.ok)
          throw new Error("ユーザーの予約状況を取得できませんでした");
        this.userReservations = await response.json();
      } catch (error) {
        alert(error.message);
      }
    },
    async createReservation() {
      try {
        // 選択された日付と時間枠の詳細を抽出
        const selectedDate = this.newReservation.date;
        const selectedSlot = this.newReservation.timeSlot;

        if (!selectedDate || !selectedSlot) {
          alert("日付と時間枠を選択してください");
          return;
        }

        // リクエストペイロードを準備
        const payload = {
          uuid: this.newReservation.uuid,
          store_id: this.newReservation.store_id,
          from: `${selectedDate}T${selectedSlot.from}:00Z`,
          to: `${selectedDate}T${selectedSlot.to}:00Z`,
        };

        const response = await fetch(
          `https://8ib4rtllwk.execute-api.ap-northeast-1.amazonaws.com/reservations`,
          {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(payload),
          }
        );

        if (response.status === 201) {
          this.reservationMessage = "予約が正常に作成されました";
        } else {
          throw new Error("予約の作成に失敗しました");
        }
      } catch (error) {
        alert(error.message);
      }
    },
  },
});

app.mount("#app");
