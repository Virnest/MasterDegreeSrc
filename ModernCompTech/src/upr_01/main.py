import dearpygui.dearpygui as dpg

if __name__ == "__main__":
    ClientHeight = 400
    ClientWidth = 300

    clBlack = (0, 0, 0, 255)
    clRed = (255, 0, 0, 255)
    clBlue = (0, 0, 255, 255)
    clGreen = (0, 255, 0, 255)
    
    dpg.create_context()

    with dpg.window(tag="Primary Window"):
        with dpg.drawlist(width=ClientWidth, height=ClientHeight):
            brush_thickness = 3
            dpg.draw_ellipse(
                    (50, 50),
                    (250, 350),
                    thickness=brush_thickness,
                    fill=clBlue,
                    color=clBlack
                )
            with dpg.draw_layer() as eye1:
                dpg.draw_ellipse(
                        (100,100),
                        (125,200),
                        thickness=brush_thickness,
                        fill=clRed,
                        color=clBlack
                    )
                dpg.draw_ellipse(
                        (105, 150),
                        (120, 175),
                        thickness=brush_thickness,
                        fill=clBlack,
                        color=clBlack
                    )
            with dpg.draw_layer() as eye2:
                dpg.draw_ellipse(
                        (175,100),
                        (200,200),
                        thickness=brush_thickness,
                        fill=clRed,
                        color=clBlack
                        )
                dpg.draw_ellipse(
                        (180, 150),
                        (195, 175),
                        thickness=brush_thickness,
                        fill=clBlack,
                        color=clBlack
                        )
            with dpg.draw_layer() as mouth:
                dpg.draw_rectangle(
                        (100, 275),
                        (200, 285),
                        thickness=brush_thickness,
                        fill=clGreen,
                        color=clBlack,
                        )
            with dpg.draw_layer() as nose:
                dpg.draw_rectangle(
                        (145, 100),
                        (155, 265),
                        thickness=brush_thickness,
                        fill=clGreen,
                        color=clBlack
                        )
            
    dpg.create_viewport(title='Маска', width=ClientWidth, height=ClientHeight)
    
    dpg.setup_dearpygui()
    dpg.show_viewport()
    dpg.set_primary_window("Primary Window", True)
    dpg.start_dearpygui()
    dpg.destroy_context()
