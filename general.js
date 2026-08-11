const modalsAttach = () => {
  for (const trigger of document.querySelectorAll("[data-modal]")) {
    const modal = document.getElementById(trigger.dataset.modal);
    if (!modal) {
      console.error("no modal found");
    }

    trigger.addEventListener("click", () => {
      if (modal.dataset.nonmodal === "true") {
        modal.show();
      } else {
        modal.showModal();
      }
    });
  }
};

document.addEventListener("DOMContentLoaded", () => {
  modalsAttach();
});
